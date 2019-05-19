package db

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	log "github.com/sirupsen/logrus"
)

var (
	dynamosession   *session.Session
	EventsTableName string
)

/**
* Initialize package variables
 */
func init() {
	EventsTableName = os.Getenv("DYNAMO_EVENTS_TABLE_NAME")
	if EventsTableName == "" {
		EventsTableName = "events"
		log.Warning("Events table name missing from environment, defaulting to ", EventsTableName)
	}
}

/**
* Get a dynamoDB connection
 */
func GetDb() *dynamodb.DynamoDB {

	if dynamosession == nil {
		log.Info("Acquiring database connection…")

		dynamoURL := os.Getenv("DYNAMO_URL")
		if dynamoURL == "" {
			dynamoURL = "http://0.0.0.0:8242"
			log.Warn("WARNING: no dynamo_url in the environment. Defaulting to ", dynamoURL)
		}

		dynamoRegion := os.Getenv("DYNAMO_REGION")
		if dynamoRegion == "" {
			dynamoRegion = "us-east-1"
			log.Warn("WARNING: no dynamo_region in the environment. Defaulting to ", dynamoURL)
		}

		config := &aws.Config{
			Region:   aws.String(dynamoRegion),
			Endpoint: aws.String(dynamoURL),
		}

		sess := session.Must(session.NewSession(config))
		dynamosession = sess

		log.Info("Successfully connected to ", dynamoURL, " in ", dynamoRegion, " region")

	}
	// fmt.Println(reflect.TypeOf(svc))

	svc := dynamodb.New(dynamosession)
	return svc
}

/**
* Checks to make sure a DynamoDB table exists
 */
func CheckTableExists(tableName string) bool {
	svc := GetDb()
	input := &dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	}
	_, err := svc.DescribeTable(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeResourceNotFoundException:
				break // this is expected and means table doesn't exist
			default:
				log.Error(err.Error()) // something unexpected happened
			}
		}
	} else {
		return true
	}
	return false

}

func provisionedCapacity() (int64, int64) {

	readCapacity := os.Getenv("DYNAMO_READ_CAPACITY")
	if readCapacity == "" {
		readCapacity = "100" // default
	}

	writeCapacity := os.Getenv("DYNAMO_WRITE_CAPACITY")
	if writeCapacity == "" {
		writeCapacity = "100" // default
	}

	iRead, err := strconv.ParseInt(readCapacity, 10, 64)
	if err != nil {
		iRead = 100 // default
	}

	iWrite, err := strconv.ParseInt(writeCapacity, 10, 64)
	if err != nil {
		iWrite = 100 // default
	}

	return iRead, iWrite

}

/**
* SaveEvent() saves a single event to dynamodb.
*/
func SaveEvent(event interface{}) {
	av, err := dynamodbattribute.MarshalMap(event)
	if err != nil {
		log.Info("Got error marshalling new event item: ", err.Error())
	}

	input := &dynamodb.PutItemInput{
        Item:      av,
        TableName: aws.String(EventsTableName),
    }

	svc := GetDb()

    _, err = svc.PutItem(input)
    if err != nil {
        fmt.Println("Got error calling PutItem:")
        fmt.Println(err.Error())
        os.Exit(1)
	}
}


/**
* CreateEventsTable creatses events table if it doesn't exist. Returns "false" if
* an error occurs
 */
func CreateEventsTable() {
	tableName := EventsTableName
	svc := GetDb()

	readCapacity, writeCapacity := provisionedCapacity()

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("eventId"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("eventType"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("eventId"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(readCapacity),
			WriteCapacityUnits: aws.Int64(writeCapacity),
		},
		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			{
				IndexName: aws.String("EventTypeIndex"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("eventType"),
						KeyType:       aws.String("HASH"),
					},
					{
						AttributeName: aws.String("eventId"),
						KeyType:       aws.String("RANGE"),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String("ALL"),
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(readCapacity),
					WriteCapacityUnits: aws.Int64(writeCapacity),
				},
			},
		},
		TableName: aws.String(tableName),
	}

	_, err := svc.CreateTable(input)
	if err != nil {
		fmt.Println("Got error calling CreateTable:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Created the table", tableName)
}

func GetAllTables() []string {

	svc := GetDb()

	input := &dynamodb.ListTablesInput{}

	tables := []string{}

	for {
		// Get the list of tables
		result, err := svc.ListTables(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case dynamodb.ErrCodeInternalServerError:
					log.Error(dynamodb.ErrCodeInternalServerError, aerr.Error())
				default:
					log.Error(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				log.Error(err.Error())
			}
			return tables
		}

		for _, n := range result.TableNames {
			tables = append(tables, *n)
		}

		// assign the last read tablename as the start for our next call to the ListTables function
		// the maximum number of table names returned in a call is 100 (default), which requires us to make
		// multiple calls to the ListTables function to retrieve all table names
		input.ExclusiveStartTableName = result.LastEvaluatedTableName

		if result.LastEvaluatedTableName == nil {
			break
		}
	}

	return tables
}
