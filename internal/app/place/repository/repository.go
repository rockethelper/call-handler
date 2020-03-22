package repository

import (
	"bytes"
	"encoding/csv"
	"errors"
	"io"

	"github.com/rockethelper/call-handler/internal/app/model"
	"github.com/rockethelper/call-handler/internal/pkg/data"
)

type Repository struct{}

func New() *Repository {
	return &Repository{}
}

func (r Repository) GermanAddressInformationList() ([]model.GermanAddressInformation, error) {
	var addresses []model.GermanAddressInformation

	zipCodeDataCSV, err := data.Asset("data/german_zip_codes.csv")
	if err != nil {
		return addresses, err
	}

	reader := csv.NewReader(bytes.NewReader(zipCodeDataCSV))
	reader.Comma = ';'
	index := 0

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if index == 0 {
			index++
			continue
		} else if err != nil {
			return addresses, err
		}

		addresses = append(addresses, model.GermanAddressInformation{
			Place:         line[0],
			PlaceAddition: line[1],
			PhoneAreaCode: line[3],
			State:         line[4],
			ZipCode:       line[2],
		})
	}

	return addresses, nil
}

func (r *Repository) FindMatchingGermanAddressInformationFor(attr string, searchValue string) (model.GermanAddressInformation, error) {
	addressInformation := model.GermanAddressInformation{}
	addressInformationList, err := r.GermanAddressInformationList()
	if err != nil {
		return addressInformation, err
	}

	matchFound := false
	for _, a := range addressInformationList {
		if a.AttributeEquals(attr, searchValue) {
			addressInformation = a
			matchFound = true
			break
		}
	}

	if matchFound {
		return addressInformation, nil
	} else {
		errMsg := "No matching address information found for attribute '" + attr + "' with value '" + searchValue + "'"
		return addressInformation, errors.New(errMsg)
	}
}
