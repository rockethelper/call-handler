package lambda

import (
	"context"
	"fmt"
	"os"

	awsLambda "github.com/aws/aws-lambda-go/lambda"
	helpRequestRepo "github.com/rockethelper/call-handler/internal/app/helprequest/repository"
	"github.com/rockethelper/call-handler/internal/app/model"
	placeRepo "github.com/rockethelper/call-handler/internal/app/place/repository"
	placeServ "github.com/rockethelper/call-handler/internal/app/place/service"
	workflowServ "github.com/rockethelper/call-handler/internal/app/workflow/service"
	"github.com/rockethelper/call-handler/internal/pkg/database"
)

func handler(ctx context.Context, input model.CallInput) (model.CallWorkflowResponse, error) {
	fmt.Println(input)

	emptyResponse := model.CallWorkflowResponse{}

	dbSession, err := database.NewSession(ctx, os.Getenv("AWS_REGION"))
	if err != nil {
		return emptyResponse, err
	}

	helpRequestRepository := helpRequestRepo.New(dbSession)

	placeRepository := placeRepo.New()
	placeService := placeServ.New(placeRepository)

	workflow := model.NewWorkflow()
	workflow.Input = input

	workflowService := workflowServ.New(workflow, helpRequestRepository, placeService)
	return workflowService.Run()
}

func RunHandler() {
	awsLambda.Start(handler)
}
