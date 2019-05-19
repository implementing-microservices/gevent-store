package db

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	log "github.com/sirupsen/logrus"
)

func GetDb() *dynamodb.DynamoDB {

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

	svc := dynamodb.New(sess)

	log.Info("Successfully connected to ", dynamoURL, " in ", dynamoRegion, " region")
	// fmt.Println(reflect.TypeOf(svc))

	return svc
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
