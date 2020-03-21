package lambda

import (
	"context"
	"fmt"

	awsLambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/rockethelper/call-handler/internal/app/model"
)

func handler(ctx context.Context, input model.CallInput) (model.AudioResponse, error) {
	// TODO: Remove later
	fmt.Print(input)

	resp := model.AudioResponse{Message: "Hello World!",
		UserZipCode: input.Details.Parameters.UserZipCode}

	return resp, nil
}

func RunHandler() {
	awsLambda.Start(handler)
}
