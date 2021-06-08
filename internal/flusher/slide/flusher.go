// Package flusher synchronizes the associated repository with data.
package flusher

import (
	"context"
	"errors"

	"github.com/ozoncp/ocp-presentation-api/internal/model"
	repo "github.com/ozoncp/ocp-presentation-api/internal/repo/slide"
)

var ErrInvalidArgument = errors.New("invalid argument")

// Flusher is the interface that wraps the basic Flush method.
type Flusher interface {
	Flush(ctx context.Context, presentations []model.Slide) ([]model.Slide, error)
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

// Flush pushes slides into the repository.
func (f *flusher) Flush(ctx context.Context, slides []model.Slide) ([]model.Slide, error) {
	if f.chunkSize == 0 {
		return nil, ErrInvalidArgument
	}

	for i, n := uint(0), uint(len(slides)); i < n; i += f.chunkSize {
		end := i + f.chunkSize
		if end > n {
			end = n
		}
		if err := f.repo.AddSlides(ctx, slides[i:end]); err != nil {
			return slides[i:], err
		}
	}

	return nil, nil
}
