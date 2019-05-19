package events

import (
	"app/db"
	"encoding/json"
	"sync"
	
	log "github.com/sirupsen/logrus"
)

/**
* Get events
 */
func GetEvents(eventType string, since string) string {

	// svc := db.GetDb()

	allTables := db.GetAllTables()

	log.Info(allTables)

	return "this is db"
}

func SaveEvents(eventType string, payload []byte) string {

	events := make([](map[string]interface{}), 0)
	json.Unmarshal(payload, &events)

	var wg sync.WaitGroup
	wg.Add(len(events))
	
	for _, event := range events { 
		event["eventType"] = eventType
		//log.Info("event : ", event)
		log.Info("Saving eventId : ", event["eventId"])
		go db.SaveEvent(event, &wg)
	}

	return "this is db response after save"
}
