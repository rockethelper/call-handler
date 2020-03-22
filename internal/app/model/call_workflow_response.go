package model

type CallWorkflowResponse struct {
	Message     string `json:"Message"`
	UserZipCode string `json:"UserZipCode"`
	ResultState string `json:"ResultState"`
}
