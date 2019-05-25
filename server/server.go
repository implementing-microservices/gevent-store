package server

import (
	"app/db"
	"app/events"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"bytes"

	gin "github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/google/uuid"
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

	errors := validateGetEvents(count, since, eventType)
	if len(errors) > 0 {
		respondWithErrors(c, errors)
		return
	}

	dbEvents := events.GetEvents(eventType, since, count)

	c.JSON(http.StatusOK, dbEvents)
}

/**
* Returns true if validation passes, false otherwise
*/
func validateGetEvents(count int64, since string , eventType string) []string {
	errors := []string{}
	
	if (count < 1){
		newErr := "The `count` parameter must be a positive integer, greater than 0"
		errors = append(errors, newErr)
	}

	_, err := uuid.Parse(since)
    if (err != nil) {
		newErr := "The `since` parameter must be a valid UUID string"
		errors = append(errors, newErr)
	}

	if (len(eventType) < 2) {
		newErr := "The `eventType` must be a valid string representing even type. Min 2 characters."
		errors = append(errors, newErr)
	}

	return errors
}

/**
* HTTP response with errors
*/
func respondWithErrors(c *gin.Context, errors []string) {
	c.Header("Content-Type", "application/json; charset=utf-8")

	var b bytes.Buffer

	for _, err := range errors {
		b.WriteString(err)
		b.WriteString(".\n ")
	}

	summaryError := b.String()
	log.Info(summaryError)
	_ = summaryError

	c.JSON(http.StatusBadRequest, gin.H{
		"errors" : errors,
		"summary":  summaryError,		
		//"status": http.StatusBadRequest,
	})
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
