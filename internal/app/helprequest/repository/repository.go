package repository

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/rockethelper/call-handler/internal/app/model"
	"github.com/rockethelper/call-handler/internal/pkg/database"
)

type Repository struct {
	Db *dynamodb.DynamoDB
}

func New(dbSession *dynamodb.DynamoDB) *Repository {
	return &Repository{Db: dbSession}
}

func (r Repository) Find(id string) (model.HelpRequest, error) {
	helpRequest := model.HelpRequest{}

	result, err := r.Db.GetItem(&dynamodb.GetItemInput{
		TableName: database.DynamoDBTableName(),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: &id,
			},
		},
	})
	if err != nil {
		return helpRequest, err
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &helpRequest)
	if err != nil {
		return helpRequest, err
	}

	return helpRequest, nil
}

func (r Repository) Create(helpRequest *model.HelpRequest) error {
	item, err := dynamodbattribute.MarshalMap(helpRequest)
	if err != nil {
		return err
	}

	params := &dynamodb.PutItemInput{
		Item:      item,
		TableName: database.DynamoDBTableName(),
	}

	_, err = r.Db.PutItem(params)
	if err != nil {
		return err
	}

	return nil
}
