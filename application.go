package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"app/server"
)

func main() {

	// isDevelopment := os.Getenv("APP_ENV") == "development"

	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = "3737"
		log.Warn("WARNING: no server port supplied in the environment. Defaulting to ", serverPort)
	}

	log.Info("Starting microservice on internal port: ", serverPort)
	server.StartServer(serverPort)
}

// @see: https://golang.org/doc/code.html
