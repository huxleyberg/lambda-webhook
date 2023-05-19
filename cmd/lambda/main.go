package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/handlerfunc"
	"github.com/huxleyberg/lambda-webhook/internal/app"
)

type awsEventHandlerFunc func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

func main() {
	h := getAWSLambdaEventHandler()
	lambda.Start(h)
}

func getAWSLambdaEventHandler() awsEventHandlerFunc {

	return func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

		a := app.New()

		lambdaProxyAdapter := handlerfunc.New(a.Handler())
		return lambdaProxyAdapter.ProxyWithContext(ctx, req)
	}
}
