package main

import (
    "context"
    "errors"
    "fmt"
    "log"
    "os"
    "strings"

    "github.com/aws/aws-lambda-go/lambda"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// GetConnectionEvent represents the event structure expected by the Lambda function
type GetConnectionEvent struct {
    UserId    string `json:"userId"`
    ContactId string `json:"contactId"`
}

// Connection represents a user's contact information
type Connection struct {
    UserId           string  `json:"userId" dynamodbav:"UserId"`
    ContactId        string  `json:"contactId" dynamodbav:"ContactId"`
    Name             string  `json:"name" dynamodbav:"Name"`
    Birthday         *string `json:"birthday,omitempty" dynamodbav:"Birthday,omitempty"`
    CheckInFrequency string  `json:"checkInFrequency" dynamodbav:"CheckInFrequency"`
    CheckInDate      string  `json:"checkInDate" dynamodbav:"CheckInDate"`
}

// handleRequest is the Lambda function handler
func handleRequest(ctx context.Context, event GetConnectionEvent) (Connection, error) {
    // Retrieve and trim environment variables for configuration
    tableName := strings.TrimSpace(os.Getenv("DYNAMODB_TABLE"))
    region := os.Getenv("AWS_REGION") // AWS_REGION is automatically set by Lambda

    log.Printf("Environment Variables - Table: '%s', Region: '%s'", tableName, region)

    // Validate environment variables
    if tableName == "" {
        log.Printf("Environment variable DYNAMODB_TABLE must be set")
        return Connection{}, errors.New("environment variable DYNAMODB_TABLE must be set")
    }

    if region == "" {
        log.Printf("AWS_REGION is not set in the environment")
        return Connection{}, errors.New("AWS_REGION is not set in the environment")
    }

    // Validate input data
    if strings.TrimSpace(event.UserId) == "" || strings.TrimSpace(event.ContactId) == "" {
        log.Printf("Invalid input: UserId and ContactId must be provided")
        return Connection{}, errors.New("invalid input: UserId and ContactId must be provided")
    }

    // Log the retrieval attempt
    log.Printf("Attempting to retrieve contact: UserId='%s', ContactId='%s'", event.UserId, event.ContactId)

    // Initialize a session to DynamoDB
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String(region),
    })
    if err != nil {
        log.Printf("Failed to create AWS session: %v", err)
        return Connection{}, err
    }
    svc := dynamodb.New(sess)

    // Define the key of the item to retrieve
    key := map[string]*dynamodb.AttributeValue{
        "UserId": {
            S: aws.String(event.UserId),
        },
        "ContactId": {
            S: aws.String(event.ContactId),
        },
    }

    // Create input for GetItem operation
    input := &dynamodb.GetItemInput{
        TableName: aws.String(tableName),
        Key:       key,
    }

    // Perform the GetItem operation
    result, err := svc.GetItem(input)
    if err != nil {
        log.Printf("Error calling GetItem: %v", err)
        return Connection{}, fmt.Errorf("failed to get item from DynamoDB: %v", err)
    }

    // Check if the item was found
    if result.Item == nil {
        log.Printf("No contact found with UserId='%s' and ContactId='%s'", event.UserId, event.ContactId)
        return Connection{}, fmt.Errorf("no contact found with UserId='%s' and ContactId='%s'", event.UserId, event.ContactId)
    }

    // Unmarshal the result into the Connection struct
    var connection Connection
    err = dynamodbattribute.UnmarshalMap(result.Item, &connection)
    if err != nil {
        log.Printf("Error unmarshalling DynamoDB item: %v", err)
        return Connection{}, fmt.Errorf("failed to unmarshal DynamoDB item: %v", err)
    }

    log.Printf("Successfully retrieved contact: %+v", connection)

    return connection, nil
}

func main() {
    // Start the Lambda function handler
    lambda.Start(handleRequest)
}
