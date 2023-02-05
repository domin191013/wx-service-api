package http

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

func StartServer(portNum int) {
	http.HandleFunc("/weather", ForecastHandler)
	port := strconv.Itoa(portNum)

	err := http.ListenAndServe("localhost:"+port, nil)
	if err != nil {
		log.Printf("the HTTP server failed to start: %s", err)
		os.Exit(1)
	}
}
