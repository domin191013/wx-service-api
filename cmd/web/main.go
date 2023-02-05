package main

import (
	"log"

	"github.com/domin191013/wx-service-api/pkg/http"
)

func main() {
	log.Print("Server started at port 3000...")

	http.StartServer(3000)
}
