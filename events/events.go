package events

import (
	"app/db"
	"encoding/json"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/google/uuid"
)

/**
* Get events
 */
func GetEvents(eventType string, since string) []interface{} {

	events := db.GetEvents(eventType, since)

	//log.Info ("db responded with events: ", events)

	return events
}

func SaveEvents(eventType string, payload []byte) string {

	events := make([](map[string]interface{}), 0)
	json.Unmarshal(payload, &events)

	var wg sync.WaitGroup
	wg.Add(len(events))
	
	for _, event := range events { 
		
		event["eventType"] = eventType
		
		u1 := uuid.Must(uuid.NewUUID())
		event["seqId"] = u1.String()

		//log.Info("event : ", event)
		log.Info("Saving eventId : ", event["eventId"])
		go db.SaveEvent(event, &wg)
	}

	wg.Wait()

	return "this is db response after save"
}
