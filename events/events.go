package events

import (
	"app/db"
	"encoding/json"

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

func SaveEvents(eventType string, payload []byte) string {

	events := make([](map[string]interface{}), 0)
	json.Unmarshal(payload, &events)

	for _, event := range events { 
		event["eventType"] = eventType
		//log.Info("event : ", event)
		log.Info("Saving eventId : ", event["eventId"])
		db.SaveEvent(event)
	}

	return "this is db response after save"
}
