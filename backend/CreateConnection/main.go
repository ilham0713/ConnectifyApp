package main

import (
    "context"
    "errors"
    "fmt"
    "log"
    "math/rand"
    "os"
    "time"

    "github.com/aws/aws-lambda-go/lambda"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Connection represents a user's contact information
type Connection struct {
    UserId           string  `json:"userId" dynamodbav:"UserId"`
    ContactId        string  `json:"contactId" dynamodbav:"ContactId"`
    Name             string  `json:"name" dynamodbav:"Name"`
    Birthday         *string `json:"birthday,omitempty" dynamodbav:"Birthday,omitempty"`
    CheckInFrequency string  `json:"checkInFrequency" dynamodbav:"CheckInFrequency"`
    CheckInDate      string  `json:"checkInDate" dynamodbav:"CheckInDate"`
}

// CreateConnectionEvent represents the event structure
type CreateConnectionEvent struct {
    UserId           string  `json:"userId"`
    ContactId        string  `json:"contactId"`
    Name             string  `json:"name"`
    Birthday         *string `json:"birthday,omitempty"`
    CheckInFrequency string  `json:"checkInFrequency"`
}

// Map of frequency to days
var frequencyToDays = map[string]int{
    "Twice a Month": 15,
    "Monthly":        30,
    "Quarterly":      90,
    "Semiannually":  180,
    "Twice a Year":   180,
}

// calculateRandomCheckInDate calculates a randomized CheckInDate based on frequency
func calculateRandomCheckInDate(frequency string) (string, error) {
    days, exists := frequencyToDays[frequency]
    if !exists {
        return "", fmt.Errorf("invalid check-in frequency: %s", frequency)
    }

    // Calculate the target date
    targetDate := time.Now().AddDate(0, 0, days)

    // Generate a random number between -15 and +15
    rand.Seed(time.Now().UnixNano())
    offset := rand.Intn(31) - 15 // Generates a number between -15 and +15

    // Apply the offset
    randomizedDate := targetDate.AddDate(0, 0, offset)

    // Format the date as YYYY-MM-DD
    return randomizedDate.Format("2006-01-02"), nil
}

// handleRequest is the Lambda function handler
func handleRequest(ctx context.Context, event CreateConnectionEvent) (string, error) {
    // Retrieve environment variables
    tableName := os.Getenv("DYNAMODB_TABLE")
    region := os.Getenv("AWS_REGION")

    if tableName == "" || region == "" {
        log.Printf("Environment variables DYNAMODB_TABLE and AWS_REGION must be set")
        return "", errors.New("environment variables not set")
    }

    // Initialize a session to DynamoDB
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String(region),
    })
    if err != nil {
        log.Printf("Failed to create session: %v", err)
        return "", err
    }
    svc := dynamodb.New(sess)

    // Calculate the randomized CheckInDate
    checkInDate, err := calculateRandomCheckInDate(event.CheckInFrequency)
    if err != nil {
        log.Printf("Error calculating CheckInDate: %v", err)
        return "", fmt.Errorf("failed to calculate CheckInDate: %v", err)
    }

    // Create a Connection instance
    connection := Connection{
        UserId:           event.UserId,
        ContactId:        event.ContactId,
        Name:             event.Name,
        Birthday:         event.Birthday,
        CheckInFrequency: event.CheckInFrequency,
        CheckInDate:      checkInDate,
    }

    log.Printf("Attempting to add connection: %+v", connection)

    // Marshal the Connection struct to DynamoDB attribute values
    av, err := dynamodbattribute.MarshalMap(connection)
    if err != nil {
        log.Printf("Error marshalling connection: %v", err)
        return "", fmt.Errorf("failed to marshal connection: %v", err)
    }

    // Create input for PutItem
    input := &dynamodb.PutItemInput{
        TableName: aws.String(tableName),
        Item:      av,
    }

    // Put the item into DynamoDB
    _, err = svc.PutItem(input)
    if err != nil {
        log.Printf("Error calling PutItem: %v", err)
        return "", fmt.Errorf("failed to put item in DynamoDB: %v", err)
    }

    log.Printf("Successfully added contact %s for user %s with CheckInDate %s", connection.ContactId, connection.UserId, connection.CheckInDate)

    return fmt.Sprintf("Successfully added contact %s for user %s with CheckInDate %s", connection.ContactId, connection.UserId, connection.CheckInDate), nil
}

func main() {
    // Make the Lambda function available to AWS Lambda
    lambda.Start(handleRequest)
}
