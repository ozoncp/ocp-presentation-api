// Package repo saves data into the database.
package repo

import (
	"context"

	"github.com/ozoncp/ocp-presentation-api/internal/model"
)

// Repo is the interface that wraps the basic methods of the database.
type Repo interface {
	AddSlide(ctx context.Context, slide *model.Slide) (uint64, error)
	AddSlides(ctx context.Context, slides []model.Slide) error
	RemoveSlide(ctx context.Context, slideID uint64) error
	GetSlide(ctx context.Context, slideID uint64) (*model.Slide, error)
}
