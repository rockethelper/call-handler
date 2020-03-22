package database

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func NewSession(awsRegion string) (*dynamodb.DynamoDB, error) {
	if session, err := session.NewSession(&aws.Config{
		Region: &awsRegion,
	}); err != nil {
		return nil, err
	} else {
		svc := dynamodb.New(session)
		return svc, nil
	}
}

func DynamoDBTableName() *string {
	return aws.String(os.Getenv("IAM_TABLE_NAME"))
}
