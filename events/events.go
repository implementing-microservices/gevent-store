package events

import (
	"app/db"

	log "github.com/sirupsen/logrus"
)

/**
* Get events
 */
func GetEvents(eventType string) string {

	// svc := db.GetDb()

	allTables := db.GetAllTables()

	log.Info(allTables)

	return "this is db"
}
