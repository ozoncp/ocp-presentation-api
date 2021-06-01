// Package repo saves data into the database.
package repo

import (
	"github.com/ozoncp/ocp-presentation-api/internal/model"
)

// Repo is the interface that wraps the basic methods of the database.
type Repo interface {
	AddSlide(slide *model.Slide) (uint64, error)
	AddSlides(slides []model.Slide) error
	RemoveSlide(slideID uint64) error
	GetSlide(slideID uint64) (*model.Slide, error)
}
