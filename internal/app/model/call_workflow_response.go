package model

type CallWorkflowResponse struct {
	Message         string `json:"Message"`
	UserZipCode     string `json:"UserZipCode"`
	UserPhoneNumber string `json:"UserPhoneNumber"`
	ResultState     string `json:"ResultState"`
}
