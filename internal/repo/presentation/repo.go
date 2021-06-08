// Package repo saves data into the database.
package repo

import (
	"context"

	"github.com/ozoncp/ocp-presentation-api/internal/model"
)

// Repo is the interface that wraps the basic methods of the database.
type Repo interface {
	AddPresentation(ctx context.Context, presentation *model.Presentation) (uint64, error)
	AddPresentations(ctx context.Context, presentations []model.Presentation) error
	RemovePresentation(ctx context.Context, presentationID uint64) error
	GetPresentation(ctx context.Context, presentationID uint64) (*model.Presentation, error)
}
