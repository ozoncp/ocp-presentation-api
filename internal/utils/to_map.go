package utils

import (
	"errors"

	"github.com/ozoncp/ocp-presentation-api/internal/models"
)

func PresentationsToMap(presentations []models.Presentation) (map[uint64]models.Presentation, error) {
	if presentations == nil {
		return nil, nil
	}

	result := make(map[uint64]models.Presentation, len(presentations))
	for i := range presentations {
		id := presentations[i].ID
		if _, found := result[id]; found {
			return nil, errors.New("identifiers are not unique")
		}
		result[id] = presentations[i]
	}

	return result, nil
}
