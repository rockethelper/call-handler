package model

import (
	"reflect"
	"strings"
)

type GermanAddressInformation struct {
	Place         string `json:"place"`
	PlaceAddition string `json:"placeAddition"`
	PhoneAreaCode string `json:"phoneAreaCode"`
	State         string `json:"state"`
	ZipCode       string `json:"zipCode"`
}

func (g GermanAddressInformation) FullPlaceName() string {
	return strings.Join([]string{g.Place, g.PlaceAddition}, " ")
}

func (g GermanAddressInformation) AttributeEquals(attr string, searchValue string) bool {
	r := reflect.ValueOf(g)
	f := reflect.Indirect(r).FieldByName(attr)
	value := string(f.String())

	return value == searchValue
}
