package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type AudiResponse struct {
	Message string `json:"Message"`
}

func Handler(ctx context.Context) (AudiResponse, error) {
	return AudiResponse{Message: "Hello World!"}, nil
}

func main() {
	lambda.Start(Handler)
}
