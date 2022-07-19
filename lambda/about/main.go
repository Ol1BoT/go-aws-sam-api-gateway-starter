package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

type body struct{ Message string }

// Handler is the Lambda function handler
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Lambda request", request.RequestContext.RequestID)
	b, _ := json.Marshal(body{Message: "about"})
	return events.APIGatewayProxyResponse{Body: string(b), StatusCode: 200}, nil
}
func main() { lambda.Start(Handler) }
