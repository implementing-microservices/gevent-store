package server

import (
	"app/db"
	"app/events"
	"fmt"
	"net/http"
	"os"
	"strconv"

	gin "github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Make sure our database is ready for us
	tableName := db.EventsTableName
	if db.CheckTableExists(tableName) {
		log.Info("Events table `", tableName, "` does exist!!!")
	} else {
		log.Info("Events table `", tableName, "` is missing!!!")
		log.Info("Creating the table")
		db.CreateEventsTable()
	}
}

/**
 * StartServer is the main entrance into the server.
 */
func StartServer(serverPort string) {

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	// @app.route('/events/<event_type>', methods=['GET'])
	router.GET("/events/:event_type", getEvents)

	router.POST("/events/:event_type", postEvents)

	router.Run(":" + serverPort)

}

func getEvents(c *gin.Context) {
	eventType := c.Param("event_type")
	since := c.Query("since")
	count, _ := strconv.ParseInt(c.Query("count"), 10, 64)
	if count == 0 {
		count = 100 // default
	}
	log.Debug("count: ", count)
	dbEvents := events.GetEvents(eventType, since, count)

	c.JSON(http.StatusOK, dbEvents)
}

func postEvents(c *gin.Context) {
	rawBody, err := c.GetRawData()

	if err != nil {
		log.Error(err.Error())
	}

	eventType := c.Param("event_type")
	saveStatus := events.SaveEvents(eventType, rawBody)

	c.String(http.StatusOK, "Response from the event store is: %s", saveStatus)
}

// Simple exit if error, to avoid putting same 4 lines of code in too many places
func abortIfErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
