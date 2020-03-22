package database

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func NewSession(ctx context.Context, awsRegion string) (*dynamodb.DynamoDB, error) {
	region := os.Getenv("AWS_REGION")
	// Initialize a session
	if session, err := session.NewSession(&aws.Config{
		Region: &region,
	}); err != nil {
		return nil, err
	} else {
		// Create DynamoDB client
		svc := dynamodb.New(session)
		return svc, nil
	}
}
