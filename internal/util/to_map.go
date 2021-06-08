// Package util implement a simple internal library for Ozon Code Platform Presentation API.
package util

import (
	"errors"

	"github.com/ozoncp/ocp-presentation-api/internal/model"
)

func PresentationsToMap(presentations []model.Presentation) (map[uint64]model.Presentation, error) {
	if presentations == nil {
		return nil, nil
	}

	result := make(map[uint64]model.Presentation, len(presentations))
	for i := range presentations {
		id := presentations[i].ID
		if _, found := result[id]; found {
			return nil, errors.New("identifiers are not unique")
		}
		result[id] = presentations[i]
	}

	return result, nil
}
