// Package flusher synchronizes the associated repository with data.
package flusher

import (
	"context"
	"errors"

	"github.com/ozoncp/ocp-presentation-api/internal/ocp-slide-api/model"
	"github.com/ozoncp/ocp-presentation-api/internal/ocp-slide-api/repo"
)

var ErrInvalidArgument = errors.New("invalid argument")

// Flusher is the interface that wraps the basic Flush method.
type Flusher interface {
	Flush(ctx context.Context, slides []model.Slide) ([]model.Slide, error)
}

type flusher struct {
	chunkSize uint
	repo      repo.Repo
}

// NewFlusher returns the Flusher interface
func NewFlusher(chunkSize uint, repo repo.Repo) Flusher {
	return &flusher{
		chunkSize: chunkSize,
		repo:      repo,
	}
}

// Flush flushes slides to the repository.
func (f *flusher) Flush(ctx context.Context, slides []model.Slide) ([]model.Slide, error) {
	if f.chunkSize == 0 {
		return nil, ErrInvalidArgument
	}

	for i, n := uint(0), uint(len(slides)); i < n; i += f.chunkSize {
		end := i + f.chunkSize
		if end > n {
			end = n
		}

		if _, err := f.repo.MultiCreateSlides(ctx, slides[i:end]); err != nil {
			return slides[i:], err
		}
	}

	return nil, nil
}
