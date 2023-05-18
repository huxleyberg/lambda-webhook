package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type RequestPayload struct {
	Type                string        `json:"type"`
	AccountID           int           `json:"account_id"`
	ID                  string        `json:"id"`
	Time                string        `json:"time"`
	ZendeskEventVersion string        `json:"zendesk_event_version"`
	Subject             string        `json:"subject"`
	Detail              RequestDetail `json:"detail"`
	Event               RequestEvent  `json:"event"`
}

type RequestDetail struct {
	CreatedAt      string `json:"created_at"`
	Email          string `json:"email"`
	ExternalID     string `json:"external_id"`
	DefaultGroupID string `json:"default_group_id"`
	ID             string `json:"id"`
	OrganizationID string `json:"organization_id"`
	Role           string `json:"role"`
	UpdatedAt      string `json:"updated_at"`
}

type RequestEvent struct {
	Current  bool `json:"current"`
	Previous bool `json:"previous"`
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Extract the Basic authentication credentials from the headers
	// username, password, err := extractBasicAuthCredentials(request.Headers)
	// if err != nil {
	// 	return events.APIGatewayProxyResponse{}, err
	// }

	// // Parse the request payload
	// var requestPayload RequestPayload
	// if err := json.Unmarshal([]byte(request.Body), &requestPayload); err != nil {
	// 	return events.APIGatewayProxyResponse{}, fmt.Errorf("failed to parse request payload: %v", err)
	// }

	// // Perform your desired logic with the request payload and authentication credentials
	// fmt.Println("Received request:", requestPayload)
	// fmt.Println("Username:", username)
	// fmt.Println("Password:", password)

	// Return a success response
	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Request processed successfully",
	}

	return response, nil
}

func extractBasicAuthCredentials(headers map[string]string) (string, string, error) {
	authorization, ok := headers["Authorization"]
	if !ok {
		return "", "", fmt.Errorf("missing Authorization header")
	}

	authParts := strings.SplitN(authorization, " ", 2)
	if len(authParts) != 2 || authParts[0] != "Basic" {
		return "", "", fmt.Errorf("invalid Authorization header format")
	}

	credentials, err := base64.StdEncoding.DecodeString(authParts[1])
	if err != nil {
		return "", "", fmt.Errorf("failed to decode Authorization credentials")
	}

	credentialsParts := strings.SplitN(string(credentials), ":", 2)
	if len(credentialsParts) != 2 {
		return "", "", fmt.Errorf("invalid Authorization credentials format")
	}

	return credentialsParts[0], credentialsParts[1], nil
}

func main() {
	lambda.Start(Handler)
}
