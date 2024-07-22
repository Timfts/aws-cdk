package main

import (
	"fmt"
	"lambda-func/app"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Username string `json:"username"`
}

func HandleRequest(event MyEvent) (string, error) {
	if event.Username == "" {
		return "", fmt.Errorf("username cannot be empty")
	}

	return fmt.Sprintf("Successfully called by - %s", event.Username), nil
}

func main() {
	myApp := app.NewApp()
	lambda.Start(func(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		switch req.Path {
		case "/register":
			return myApp.ApiHandler.RegisterUserHandler(req)
		case "/login":
			return myApp.ApiHandler.LoginUser(req)
		default:
			return events.APIGatewayProxyResponse{
				Body:       "not found",
				StatusCode: http.StatusNotFound,
			}, nil
		}
	})
}
