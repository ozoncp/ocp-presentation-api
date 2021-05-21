package utils

import "github.com/ozoncp/ocp-presentation-api/internal/models"

// SplitPresentationsToBulks splits a slice into uniform chunks.
func SplitPresentationsToBulks(presentations []models.Presentation, butchSize uint) [][]models.Presentation {
	var bulks [][]models.Presentation

	if butchSize == 0 {
		return bulks
	}

	for i, n := uint(0), uint(len(presentations)); i < n; i += butchSize {
		end := i + butchSize
		if end > n {
			end = n
		}
		bulks = append(bulks, presentations[i:end])
	}

	return bulks
}
