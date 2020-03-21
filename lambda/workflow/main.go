package main

import (
	lambdaWorkflow "github.com/rockethelper/call-handler/internal/app/workflow/delivery/lambda"
)

func main() {
	lambdaWorkflow.RunHandler()
}
