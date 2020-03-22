package model

type CallInputParameters struct {
	Action          string `json:"Action"`
	UserPhoneNumber string `json:"UserPhoneNumber"`
	UserZipCode     string `json:"UserZipCode"`
}

type CallInputDetails struct {
	Parameters CallInputParameters `json:"Parameters"`
}

type CallInput struct {
	Name    string           `json:"Name"`
	Details CallInputDetails `json:"Details"`
}
