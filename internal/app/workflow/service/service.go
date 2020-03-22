package service

import (
	"errors"

	helpRequestRepo "github.com/rockethelper/call-handler/internal/app/helprequest/repository"
	"github.com/rockethelper/call-handler/internal/app/model"
	placeServ "github.com/rockethelper/call-handler/internal/app/place/service"
)

type Service struct {
	HelpRequestRepository *helpRequestRepo.Repository
	PlaceService          *placeServ.Service
	Workflow              *model.Workflow
}

func New(workflow *model.Workflow, helpRequestRepository *helpRequestRepo.Repository, placeService *placeServ.Service) *Service {
	return &Service{
		HelpRequestRepository: helpRequestRepository,
		PlaceService:          placeService,
		Workflow:              workflow,
	}
}

func (s Service) Run() (model.CallWorkflowResponse, error) {
	response := model.CallWorkflowResponse{}

	workflow := s.Workflow
	if workflow.Action == "" {
		workflow.Action = workflow.Input.Details.Parameters.Action
	}

	switch action := workflow.Action; action {
	case "checkIfHelpRequestForUserCanBeCreated":
		return s.checkIfHelpRequestForUserCanBeCreated()
	case "createUserHelpRequest":
		return s.createUserHelpRequest()
	case "validateUserZipCodeInput":
		return s.validateUserZipCodeInput()
	default:
		response.ResultState = "fail"
		return response, errors.New("Unsupported workflow action")
	}
}

func (s Service) checkIfHelpRequestForUserCanBeCreated() (model.CallWorkflowResponse, error) {
	userZipCode := s.Workflow.Input.Details.Parameters.UserZipCode
	helpRequestType := s.Workflow.Input.Details.Parameters.HelpRequestType
	response := model.CallWorkflowResponse{}

	if helpRequestType == "question-help" {
		response.Message = "Request could be created"
		response.ResultState = "success"

		return response, nil
	} else {
		helpRequest := model.NewHelpRequest()
		helpRequest.RequestType = helpRequestType
		helpRequest.PhoneNumber = s.Workflow.Input.Details.ContactData.CustomerEndpoint.Address
		helpRequest.ZipCode = userZipCode
		err := helpRequest.GenerateID()
		if err != nil {
			response.ResultState = "fail"
			return response, err
		}

		result, err := s.HelpRequestRepository.Find(helpRequest.ID)
		// No entry found
		if err != nil {
			response.Message = "Request could be created"
			response.UserZipCode = userZipCode
			response.UserPhoneNumber = helpRequest.PhoneNumber
			response.ResultState = "success"

			return response, nil
		}

		if result.SecondsSinceLastUpdate() > 900 {
			response.Message = "Request could be created"
			response.UserZipCode = userZipCode
			response.UserPhoneNumber = helpRequest.PhoneNumber
			response.ResultState = "success"

			return response, nil
		} else {
			response.ResultState = "fail"

			return response, errors.New("Last HelpRequest for this user and type was created less than 900 seoncds ago.")
		}
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
	err := helpRequest.GenerateID()
	if err != nil {
		response.ResultState = "fail"
		return response, err
	}

	err = s.HelpRequestRepository.Create(helpRequest)
	if err != nil {
		response.ResultState = "fail"
		return response, err
	}

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
