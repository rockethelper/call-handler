package main

import (
	"log"

	placeRepo "github.com/rockethelper/call-handler/internal/app/place/repository"
)

func main() {
	placeRepository := placeRepo.New()

	_, err := placeRepository.FindMatchingGermanAddressInformationFor("ZipCode", "99098")
	if err != nil {
		log.Fatal(err)
	}
}
