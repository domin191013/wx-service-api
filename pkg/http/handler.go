package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/domin191013/wx-service-api/config"
	"github.com/domin191013/wx-service-api/pkg/types"
	"github.com/domin191013/wx-service-api/pkg/utils"
)

func ForecastHandler(w http.ResponseWriter, r *http.Request) {
	var lattitude string
	var longitude string

	if lat := r.URL.Query().Get("lat"); len(lat) > 0 {
		lattitude = lat
	}
	if lon := r.URL.Query().Get("lon"); len(lon) > 0 {
		longitude = lon
	}

	if lattitude == "" || longitude == "" {
		w.Write([]byte("Please put lattitude, longitude query params.."))
		return
	}

	forecast := utils.GetForecastFromOpenWeather(lattitude, longitude)
	temperature := config.Temps[1]
	conditions := ""
	alerts := ""

	// convert response to map format, get temperature and condition there
	if current, ok := forecast["current"].(map[string]interface{}); ok {
		currentTemp := current["temp"].(float64)
		if currentTemp > config.Hot {
			temperature = config.Temps[2]
		}
		if currentTemp < config.Cold {
			temperature = config.Temps[0]
		}

		// get weather from response, generate conditions string
		if weatherList, aok := current["weather"].([]interface{}); aok {
			if len(weatherList) > 0 {
				weather := weatherList[0].(map[string]interface{})
				conditions = fmt.Sprintf("%v: %v", weather["main"], weather["description"])
			}
		}
	}

	// get alerts response if there's any
	if alertList, ok := forecast["alerts"].([]interface{}); ok {
		if len(alertList) > 0 {
			a := alertList[0].(map[string]interface{})
			alerts = fmt.Sprintf("%v", a["event"])
		}
	}

	if json, err := json.Marshal(types.Weather{temperature, conditions, alerts}); err == nil {
		w.Write(json)
		return
	}

	w.Write([]byte("Internal Server error"))
}
