package main

import (
	"app/server"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {

	// isDevelopment := os.Getenv("APP_ENV") == "development"

	serverPort := os.Getenv("API_PORT")
	if serverPort == "" {
		serverPort = "3737"
		log.Warn("WARNING: no server port supplied in the environment. Defaulting to ", serverPort)
	}

	log.Info("Starting the event-store API microservice on internal port: ", serverPort)
	server.StartServer(serverPort)
}

// @see: https://golang.org/doc/code.html
