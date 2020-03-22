package repository

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Repository struct {
	Db *dynamodb.DynamoDB
}

func New(db *dynamodb.DynamoDB) *Repository {
	return &Repository{Db: db}
}

func (r Repository) DynamoDBTableName() *string {
	return aws.String(os.Getenv("IAM_TABLE_NAME"))
}

func (r Repository) Create() error {
	// 	item, err := dynamodbattribute.MarshalMap(helpRequest)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	params := &dynamodb.PutItemInput{
	// 		Item:      item,
	// 		TableName: r.DynamoDBTableName(),
	// 	}

	// 	_, err = dbSession.PutItem(params)
	// 	if err != nil {
	// 		return err
	// 	}

	return nil
}
