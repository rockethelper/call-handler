package lambda

import (
	"context"

	awsLambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/rockethelper/call-handler/internal/app/model"
	placeRepo "github.com/rockethelper/call-handler/internal/app/place/repository"
	placeServ "github.com/rockethelper/call-handler/internal/app/place/service"
	workflowServ "github.com/rockethelper/call-handler/internal/app/workflow/service"
)

func handler(ctx context.Context, input model.CallInput) (model.CallWorkflowResponse, error) {
	placeRepository := placeRepo.New()
	placeService := placeServ.New(placeRepository)

	workflow := model.NewWorkflow()
	workflow.Input = input

	workflowService := workflowServ.New(workflow, placeService)
	return workflowService.Run()
}

func RunHandler() {
	awsLambda.Start(handler)
}
