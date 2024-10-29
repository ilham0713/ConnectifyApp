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

// Connection represents a user's contact information
type Connection struct {
    UserId           string  `json:"userId" dynamodbav:"UserId"`
    ContactId        string  `json:"contactId" dynamodbav:"ContactId"`
    Name             string  `json:"name" dynamodbav:"Name"`
    Birthday         *string `json:"birthday,omitempty" dynamodbav:"Birthday,omitempty"`
    CheckInFrequency string  `json:"checkInFrequency" dynamodbav:"CheckInFrequency"`
    CheckInDate      string  `json:"checkInDate" dynamodbav:"CheckInDate"`
}

// UpdateConnectionEvent represents the event structure expected by the Lambda function
type UpdateConnectionEvent struct {
    UserId           string  `json:"userId" dynamodbav:"UserId"`
    ContactId        string  `json:"contactId" dynamodbav:"ContactId"`
    Name             *string `json:"name,omitempty" dynamodbav:"Name,omitempty"`
    Birthday         *string `json:"birthday,omitempty" dynamodbav:"Birthday,omitempty"`
    CheckInFrequency *string `json:"checkInFrequency,omitempty" dynamodbav:"CheckInFrequency,omitempty"`
}

// handleRequest is the Lambda function handler
func handleRequest(ctx context.Context, event UpdateConnectionEvent) (string, error) {
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

    // Validate input data
    if strings.TrimSpace(event.UserId) == "" || strings.TrimSpace(event.ContactId) == "" {
        log.Printf("Invalid input: UserId and ContactId must be provided")
        return "", errors.New("invalid input: UserId and ContactId must be provided")
    }

    // Validate CheckInFrequency if provided
    if event.CheckInFrequency != nil {
        validFrequencies := map[string]bool{
            "Twice a Month": true,
            "Monthly":        true,
            "Quarterly":      true,
            "Semiannually":  true,
            "Twice a Year":   true,
        }
        if !validFrequencies[*event.CheckInFrequency] {
            log.Printf("Invalid CheckInFrequency: %s", *event.CheckInFrequency)
            return "", fmt.Errorf("invalid CheckInFrequency: %s", *event.CheckInFrequency)
        }
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

    // Prepare the update expression
    updateExpression := "SET"
    expressionAttributeNames := make(map[string]*string)
    expressionAttributeValues := make(map[string]*dynamodb.AttributeValue)
    fieldsToUpdate := 0

    // Dynamically build the update expression based on provided fields
    if event.Name != nil {
        updateExpression += " #N = :n"
        expressionAttributeNames["#N"] = aws.String("Name")
        expressionAttributeValues[":n"] = &dynamodb.AttributeValue{S: event.Name}
        fieldsToUpdate++
    }

    if event.Birthday != nil {
        if fieldsToUpdate > 0 {
            updateExpression += ","
        }
        updateExpression += " #B = :b"
        expressionAttributeNames["#B"] = aws.String("Birthday")
        expressionAttributeValues[":b"] = &dynamodb.AttributeValue{S: event.Birthday}
        fieldsToUpdate++
    }

    if event.CheckInFrequency != nil {
        if fieldsToUpdate > 0 {
            updateExpression += ","
        }
        updateExpression += " #C = :c"
        expressionAttributeNames["#C"] = aws.String("CheckInFrequency")
        expressionAttributeValues[":c"] = &dynamodb.AttributeValue{S: event.CheckInFrequency}
        fieldsToUpdate++
    }

    if fieldsToUpdate == 0 {
        log.Printf("No fields to update")
        return "", errors.New("no fields provided to update")
    }

    // Define the key of the item to update
    key := map[string]*dynamodb.AttributeValue{
        "UserId": {
            S: aws.String(event.UserId),
        },
        "ContactId": {
            S: aws.String(event.ContactId),
        },
    }

    // Create input for UpdateItem operation
    input := &dynamodb.UpdateItemInput{
        TableName:                 aws.String(tableName),
        Key:                       key,
        UpdateExpression:          aws.String(updateExpression),
        ExpressionAttributeNames:  expressionAttributeNames,
        ExpressionAttributeValues: expressionAttributeValues,
        ReturnValues:              aws.String("UPDATED_NEW"),
    }

    // Perform the UpdateItem operation
    result, err := svc.UpdateItem(input)
    if err != nil {
        log.Printf("Failed to update item: %v", err)
        return "", fmt.Errorf("failed to update item: %v", err)
    }

    log.Printf("Successfully updated contact %s for user %s. Updated attributes: %v", event.ContactId, event.UserId, result.Attributes)

    // Initialize updatedAttributes map
    updatedAttributes := make(map[string]string)

    if result.Attributes != nil {
        for k, v := range result.Attributes {
            if v.S != nil {
                updatedAttributes[k] = *v.S
            }
        }
    }

    // Marshal the updated attributes to JSON for the response
    updatedAttributesJSON, err := dynamodbattribute.MarshalMap(updatedAttributes)
    if err != nil {
        log.Printf("Error marshalling updated attributes: %v", err)
        return "", fmt.Errorf("failed to marshal updated attributes: %v", err)
    }

    // Return a success message with updated attributes
    return fmt.Sprintf("Successfully updated contact %s for user %s. Updated attributes: %v", event.ContactId, event.UserId, updatedAttributesJSON), nil
}

func main() {
    // Start the Lambda function handler
    lambda.Start(handleRequest)
}
