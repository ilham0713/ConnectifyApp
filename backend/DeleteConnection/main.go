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
)

// DeleteContactEvent represents the event structure expected by the Lambda function
type DeleteContactEvent struct {
    UserId    string `json:"userId"`
    ContactId string `json:"contactId"`
}

// handleRequest is the Lambda function handler
func handleRequest(ctx context.Context, event DeleteContactEvent) (string, error) {
    // Retrieve and trim environment variables for configuration
    tableName := strings.TrimSpace(os.Getenv("DYNAMODB_TABLE"))
    region := os.Getenv("AWS_REGION") // AWS_REGION is automatically set by Lambda

    log.Printf("Environment Variables - Table: '%s', Region: '%s'", tableName, region)

    // Validate environment variables
    if tableName == "" {
        log.Printf("Environment variable DYNAMODB_TABLE must be set")
        return "", errors.New("environment variable DYNAMODB_TABLE must be set")
    }

    if region == "" {
        log.Printf("AWS_REGION is not set in the environment")
        return "", errors.New("AWS_REGION is not set in the environment")
    }

    // Initialize a session to DynamoDB
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String(region),
    })
    if err != nil {
        log.Printf("Failed to create AWS session: %v", err)
        return "", err
    }
    svc := dynamodb.New(sess)

    // Define the key of the item to delete
    key := map[string]*dynamodb.AttributeValue{
        "UserId": {
            S: aws.String(event.UserId),
        },
        "ContactId": {
            S: aws.String(event.ContactId),
        },
    }

    log.Printf("Attempting to delete contact: UserId=%s, ContactId=%s", event.UserId, event.ContactId)

    // Create input for DeleteItem operation
    input := &dynamodb.DeleteItemInput{
        TableName: aws.String(tableName),
        Key:       key,
    }

    // Perform the DeleteItem operation
    _, err = svc.DeleteItem(input)
    if err != nil {
        log.Printf("Error calling DeleteItem: %v", err)
        return "", fmt.Errorf("failed to delete item from DynamoDB: %v", err)
    }

    log.Printf("Successfully deleted contact %s for user %s", event.ContactId, event.UserId)

    return fmt.Sprintf("Successfully deleted contact %s for user %s", event.ContactId, event.UserId), nil
}

func main() {
    // Start the Lambda function handler
    lambda.Start(handleRequest)
}
