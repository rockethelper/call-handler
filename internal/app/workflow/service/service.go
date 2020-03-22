package service

import (
	"errors"

	"github.com/rockethelper/call-handler/internal/app/model"
	placeServ "github.com/rockethelper/call-handler/internal/app/place/service"
)

type Service struct {
	PlaceService *placeServ.Service
	Workflow     *model.Workflow
}

func New(workflow *model.Workflow, placeService *placeServ.Service) *Service {
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
	case "validateUserZipCodeInput":
		return s.validateUserZipCodeInput()
	default:
		response.ResultState = "fail"
		return response, errors.New("Unsupported workflow action")
	}
}

func (s Service) validateUserZipCodeInput() (model.CallWorkflowResponse, error) {
	userZipCode := s.Workflow.Input.Details.Parameters.UserZipCode
	response := model.CallWorkflowResponse{}

	if s.Workflow.Language != "de" {
		response.ResultState = "fail"
		return response, errors.New("Zip code check just supported for language 'de'")
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
