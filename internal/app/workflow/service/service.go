package service

import (
	"context"
	"errors"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	helpRequestRepo "github.com/rockethelper/call-handler/internal/app/helprequest/repository"
	"github.com/rockethelper/call-handler/internal/app/model"
	placeServ "github.com/rockethelper/call-handler/internal/app/place/service"
	"github.com/rockethelper/call-handler/internal/pkg/database"
)

type Service struct {
	HelpRequestRepository *helpRequestRepo.Repository
	PlaceService          *placeServ.Service
	Workflow              *model.Workflow
}

func New(workflow *model.Workflow, helpRequestRepository *helpRequestRepo.Repository, placeService *placeServ.Service) *Service {
	return &Service{
		PlaceService: placeService,
		Workflow:     workflow,
	}
}

func (s Service) Run() (model.CallWorkflowResponse, error) {
	response := model.CallWorkflowResponse{}

	workflow := s.Workflow
	if workflow.Action == "" {
		workflow.Action = workflow.Input.Details.Parameters.Action
	}

	switch action := workflow.Action; action {
	case "createUserHelpRequest":
		return s.createUserHelpRequest()
	case "validateUserZipCodeInput":
		return s.validateUserZipCodeInput()
	default:
		response.ResultState = "fail"
		return response, errors.New("Unsupported workflow action")
	}
}

func (s Service) createUserHelpRequest() (model.CallWorkflowResponse, error) {
	userZipCode := s.Workflow.Input.Details.Parameters.UserZipCode
	helpRequestType := s.Workflow.Input.Details.Parameters.HelpRequestType
	response := model.CallWorkflowResponse{}

	if helpRequestType != "direct-help" && helpRequestType != "question-help" {
		response.ResultState = "fail"
		return response, errors.New("HelpRequestType is not supported")
	}

	helpRequest := model.NewHelpRequest()
	helpRequest.RequestType = helpRequestType
	helpRequest.PhoneNumber = s.Workflow.Input.Details.ContactData.CustomerEndpoint.Address
	helpRequest.ZipCode = userZipCode
	helpRequest.GenerateID(s.Workflow.Input.Details.ContactData.ContactId)

	dbSession, err := database.NewSession(context.TODO(), os.Getenv("AWS_REGION"))
	if err != nil {
		return response, err
	}

	item, err := dynamodbattribute.MarshalMap(helpRequest)
	if err != nil {
		return response, err
	}

	params := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(os.Getenv("IAM_TABLE_NAME")),
	}

	_, err = dbSession.PutItem(params)
	if err != nil {
		return response, err
	}
	// err := s.HelpRequestRepository.Create()
	// fmt.Println(err)
	// if err != nil {
	// 	response.ResultState = "fail"
	// 	return response, errors.New("It doesn't work!")
	// }

	response.UserZipCode = userZipCode
	response.UserPhoneNumber = helpRequest.PhoneNumber
	response.ResultState = "success"

	return response, nil
}

func (s Service) validateUserZipCodeInput() (model.CallWorkflowResponse, error) {
	userZipCode := s.Workflow.Input.Details.Parameters.UserZipCode
	response := model.CallWorkflowResponse{}

	if s.Workflow.Language != "de" {
		response.ResultState = "fail"
		return response, errors.New("Zip code check is just supported for language 'de'")
	}

	_, err := s.PlaceService.FindMatchingGermanAddressInformationFor("ZipCode", userZipCode)
	if err != nil {
		response.ResultState = "fail"
		return response, errors.New("Zip Code not valid")
	}

	response.UserZipCode = userZipCode
	response.ResultState = "success"

	return response, nil
}
