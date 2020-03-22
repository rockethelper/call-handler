package model

type CallInputCustomerEndpoint struct {
	Address string `json:"Address"`
	Type    string `json:"Type"`
}

type CallInputContactData struct {
	Channel          string                    `json:"Channel"`
	ContactId        string                    `json:"ContactId"`
	CustomerEndpoint CallInputCustomerEndpoint `json:"CustomerEndpoint"`
}

type CallInputParameters struct {
	Action          string `json:"Action"`
	HelpRequestType string `json:"HelpRequestType"`
	UserPhoneNumber string `json:"UserPhoneNumber"`
	UserZipCode     string `json:"UserZipCode"`
}

type CallInputDetails struct {
	ContactData CallInputContactData `json:"ContactData"`
	Parameters  CallInputParameters  `json:"Parameters"`
}

type CallInput struct {
	Name    string           `json:"Name"`
	Details CallInputDetails `json:"Details"`
}
