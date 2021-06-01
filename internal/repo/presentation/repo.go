// Package repo saves data into the database.
package repo

import "github.com/ozoncp/ocp-presentation-api/internal/model"

// Repo is the interface that wraps the basic methods of the database.
type Repo interface {
	AddPresentation(presentation *model.Presentation) (uint64, error)
	AddPresentations(presentations []model.Presentation) error
	RemovePresentation(presentationID uint64) error
	GetPresentation(presentationID uint64) (*model.Presentation, error)
}
