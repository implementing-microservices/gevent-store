package server

import (
	"fmt"
	"net/http"
	"os"

	gin "github.com/gin-gonic/gin"
)

/**
 * StartServer is the main entrance into the server.
 */
func StartServer(serverPort string) {

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()
	router.GET("/", homeEndpoint)
	router.Run(":" + serverPort)

}

func homeEndpoint(c *gin.Context) {
	toGreetName := "Go enthusiasts"
	c.String(http.StatusOK, "Hello %s! Isn't this just awesome? Hot-reloading is great! I love it.", toGreetName)
}

// Simple exit if error, to avoid putting same 4 lines of code in too many places
func abortIfErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
