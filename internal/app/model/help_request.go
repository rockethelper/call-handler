package model

import (
	"strings"
	"time"
)

type HelpRequest struct {
	ID          string `json:"id"`
	RequestType string `json:"request_type"`
	PhoneNumber string `json:"phone_number"`
	ZipCode     string `json:"zip_code"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

func NewHelpRequest() *HelpRequest {
	return &HelpRequest{
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}
}

func (h *HelpRequest) GenerateID(callContactID string) {
	h.ID = strings.Join([]string{callContactID, h.RequestType, h.PhoneNumber}, "-")
}
