package main

import (
	"project-root/handlers"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handlers.CreateConnectionHandler)
}
