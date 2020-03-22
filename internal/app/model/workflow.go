package model

type Workflow struct {
	Action   string    `json:"action"`
	Input    CallInput `json:"input"`
	Language string    `json:"input"`
	Status   string    `json:"status"`
}

func NewWorkflow() *Workflow {
	return &Workflow{
		Language: "de",
		Status:   "init",
	}
}
