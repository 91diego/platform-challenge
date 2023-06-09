package pkg

import (
	"fmt"
	"strings"

	"github.com/91diego/platform-challenge/pkg/models"
	"github.com/91diego/platform-challenge/pkg/utils"
)

// ShipmentAssignments
func ShipmentAssignments(shipmentsFilePath, driversFilePath *string) ([]models.AssignmentResponse, error) {
	result, err := getSuitabilityScore(shipmentsFilePath, driversFilePath)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getSuitabilityScore
func getSuitabilityScore(shipmentsFile, driversFile *string) ([]models.AssignmentResponse, error) {

	shipmentAddressFile, err := utils.FileReader(shipmentsFile)
	if err != nil {
		return nil, fmt.Errorf("shipment file is not valid")
	}
	driverFile, err := utils.FileReader(driversFile)
	if err != nil {
		return nil, fmt.Errorf("driver file is not valid")
	}

	var name string
	var driverName string
	var baseSuitabilityScore float32
	var suitabilityScore float32
	var tmpPosition int // slice position
	var assignmentResponse models.AssignmentResponse

	results := []models.AssignmentResponse{}
	for _, shipment := range shipmentAddressFile {

		name = ""
		driverName = ""
		baseSuitabilityScore = 0
		suitabilityScore = 0

		if len(shipment)&1 == 1 {
			for k, d := range driverFile {
				baseSuitabilityScore, name = getBaseSSConsonants(d)
				if baseSSIncreased(len(shipment), len(name)) {
					baseSuitabilityScore = (baseSuitabilityScore * .50) + baseSuitabilityScore
				}
				if suitabilityScore < baseSuitabilityScore {
					tmpPosition = k
					driverName = name
					suitabilityScore = baseSuitabilityScore
				}
			}
		} else {
			for k, d := range driverFile {
				baseSuitabilityScore, name = getBaseSSVowels(d)
				if baseSSIncreased(len(shipment), len(name)) {
					baseSuitabilityScore = (baseSuitabilityScore * .50) + baseSuitabilityScore
				}
				if suitabilityScore < baseSuitabilityScore {
					tmpPosition = k
					driverName = name
					suitabilityScore = baseSuitabilityScore
				}
			}
		}

		driverFile = append(driverFile[:tmpPosition], driverFile[tmpPosition+1:]...)
		assignmentResponse = models.AssignmentResponse{
			ShipmentAddress: strings.TrimSpace(shipment),
			Driver: models.Driver{
				Name: strings.TrimSpace(driverName),
				SS:   suitabilityScore,
			},
		}
		results = append(results, assignmentResponse)
	}
	return results, nil
}

// getBaseSSConsonants
func getBaseSSConsonants(name string) (res float32, driverName string) {
	var c int
	var consonant rune = 197
	for consonant = 97; consonant <= 122; consonant++ {
		c = c + strings.Count(strings.ToLower(name), string(consonant))
	}
	res = float32(c)
	driverName = name
	return
}

// baseSSIncreased
func baseSSIncreased(shipment, driverName int) bool {
	return shipment == driverName
}

// getBaseSSVowels
func getBaseSSVowels(name string) (res float32, driverName string) {
	a := strings.Count(strings.ToLower(name), "a")
	e := strings.Count(strings.ToLower(name), "e")
	i := strings.Count(strings.ToLower(name), "i")
	o := strings.Count(strings.ToLower(name), "o")
	u := strings.Count(strings.ToLower(name), "u")
	res = float32(a+e+i+o+u) + 1.5
	driverName = name
	return
}
