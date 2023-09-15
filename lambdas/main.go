package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type App struct {
	id string
}

func newApp(id string) *App {
	return &App{
		id: id,
	}
}

func (app *App) Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	resBody := map[string]string{
		"message": "Testing lambda",
	}
	resJson, err := json.Marshal(resBody)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers: map[string]string{
				"Content-type": "application/json",
			},
			Body: `{
				"error":"Internal Server Error"
			}`,
		}, nil
	}

	response := events.APIGatewayProxyResponse{
		Body:       string(resJson),
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type":                           "text/plain",
			"Access-Control-Allow-Origin":            "*",
			"Access-Control-Allow-Headers":           "Content-Type",
			"Access-Control-Allow-Methods":           "OPTIONS, POST, GET",
			"Access-Control-Allow-Allow-Credentials": "true",
		},
	}
	return response, nil
}

func main() {
	id := "exampleId"
	app := newApp(id)
	lambda.Start(app.Handler)
}
