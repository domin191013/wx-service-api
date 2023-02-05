package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/domin191013/wx-service-api/config"
)

// Make a request to openweatherapi, convert to a Weather struct and send it
func GetForecastFromOpenWeather(lat, long string) map[string]interface{} {
	url := fmt.Sprintf(config.GetWeatherUrl, lat, long, config.APIID)
	req, _ := http.NewRequest("GET", url, bytes.NewReader([]byte("")))
	h := req.Header
	h.Add("Accept", "application/json")
	h.Add("Connection", "close")

	client := &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var forecast map[string]interface{}
	json.Unmarshal(body, &forecast)
	return forecast
}
