package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"project-root/models"
	"project-root/services"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type ContactRequest struct {
	UserId string `json:"userId"`
	Name string `json:"name"`
	LastCheckIn string `json:"lastCheckIn"`
	CheckInFreq string `json:"checkInFreq"`
	Birthday string `json:"birthday"`
}