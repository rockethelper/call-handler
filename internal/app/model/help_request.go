package model

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strings"
	"time"
)

type HelpRequest struct {
	ID          string    `json:"id"`
	RequestType string    `json:"request_type"`
	PhoneNumber string    `json:"phone_number"`
	ZipCode     string    `json:"zip_code"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

func NewHelpRequest() *HelpRequest {
	return &HelpRequest{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (h *HelpRequest) GenerateID() error {
	if h.RequestType == "" || h.PhoneNumber == "" || h.ZipCode == "" {
		return errors.New("RequestType, PhoneNumber and ZipCode for HelpRequest need to be present")
	}

	idStr := strings.Join([]string{h.PhoneNumber, h.ZipCode, h.RequestType}, "-")
	hash := md5.Sum([]byte(idStr))
	h.ID = "help-request-" + hex.EncodeToString(hash[:])

	return nil
}

func (h HelpRequest) SecondsSinceLastUpdate() float64 {
	return time.Since(h.UpdatedAt).Seconds()
}
