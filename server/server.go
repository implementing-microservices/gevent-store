package server

import (
	"app/events"
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

	// @app.route('/events/<event_type>', methods=['GET'])
	router.GET("/events/:event_type", getEvents)

	router.Run(":" + serverPort)

}

func getEvents(c *gin.Context) {
	eventType := c.Param("event_type")
	dbEvents := events.GetEvents(eventType)
	c.String(http.StatusOK, "Response from the event store is: %s", dbEvents)
}

// Simple exit if error, to avoid putting same 4 lines of code in too many places
func abortIfErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
