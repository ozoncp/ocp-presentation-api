// Package repo saves data into the database.
package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ozoncp/ocp-presentation-api/internal/ocp-slide-api/model"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

const tableName = "slide"

var (
	ErrSlideNotFound = errors.New("slide not found")
	ErrUnknown       = errors.New("unknown")
)

// Repo is the interface that wraps the basic methods of the database.
type Repo interface {
	CreateSlide(ctx context.Context, slide model.Slide) (uint64, error)
	MultiCreateSlides(ctx context.Context, slides []model.Slide) (int64, error)
	UpdateSlide(ctx context.Context, slide model.Slide) (bool, error)
	DescribeSlide(ctx context.Context, slideID uint64) (*model.Slide, error)
	ListSlides(ctx context.Context, limit uint64, offset uint64) ([]model.Slide, error)
	RemoveSlide(ctx context.Context, slideID uint64) (bool, error)
}

type repo struct {
	db *sqlx.DB
}

// NewRepo returns the Repo interface
func NewRepo(db *sqlx.DB) Repo {
	return &repo{
		db: db,
	}
}

// CreateSlide creates the new slide into the database and returns the ID of the created slide
func (r *repo) CreateSlide(ctx context.Context, slide model.Slide) (uint64, error) {
	query := squirrel.Insert(tableName).
		Columns(
			"presentation_id",
			"number",
			"type",
		).
		Values(
			slide.PresentationID,
			slide.Number,
			slide.Type,
		).
		Suffix("RETURNING \"id\"").
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	err := query.QueryRowContext(ctx).Scan(&slide.ID)
	if err != nil {
		return 0, err
	}

	return slide.ID, nil
}

// MultiCreateSlides creates the new slides into the database and returns the number of the created slides
func (r *repo) MultiCreateSlides(ctx context.Context, slides []model.Slide) (int64, error) {
	query := squirrel.Insert(tableName).
		Columns(
			"presentation_id",
			"number",
			"type",
		).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	for i := range slides {
		query = query.Values(
			slides[i].PresentationID,
			slides[i].Number,
			slides[i].Type,
		)
	}

	result, err := query.ExecContext(ctx)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// UpdateSlide updates the slide into the database
// and returns true, nil if the value was updated; otherwise, false and the error
func (r *repo) UpdateSlide(ctx context.Context, slide model.Slide) (bool, error) {
	query := squirrel.Update(tableName).
		Set("presentation_id", slide.PresentationID).
		Set("number", slide.Number).
		Set("type", slide.Type).
		Where(squirrel.Eq{"id": slide.ID}).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	result, err := query.ExecContext(ctx)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if rowsAffected <= 0 {
		return false, ErrSlideNotFound
	}

	return true, nil
}

// DescribeSlide describes the slide from the database by the slide ID
func (r *repo) DescribeSlide(ctx context.Context, slideID uint64) (*model.Slide, error) {
	query := squirrel.Select(
		"id",
		"presentation_id",
		"number",
		"type",
	).
		From(tableName).
		Where(squirrel.Eq{"id": slideID}).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	rows, err := query.QueryContext(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrSlideNotFound
		}

		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	var slides []model.Slide
	for rows.Next() {
		var slide model.Slide
		if err = rows.Scan(
			&slide.ID,
			&slide.PresentationID,
			&slide.Number,
			&slide.Type); err != nil {
			return nil, err
		}
		slides = append(slides, slide)
	}

	if len(slides) == 0 {
		return nil, ErrUnknown
	}

	return &slides[0], nil
}

func (r *repo) ListSlides(ctx context.Context, limit uint64, offset uint64) ([]model.Slide, error) {
	query := squirrel.Select(
		"id",
		"presentation_id",
		"number",
		"type",
	).
		From(tableName).
		RunWith(r.db).
		Limit(limit).
		Offset(offset).
		PlaceholderFormat(squirrel.Dollar)

	rows, err := query.QueryContext(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrSlideNotFound
		}

		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	var slides []model.Slide
	for rows.Next() {
		var slide model.Slide
		if err = rows.Scan(
			&slide.ID,
			&slide.PresentationID,
			&slide.Number,
			&slide.Type); err != nil {
			return nil, err
		}
		slides = append(slides, slide)
	}
	return slides, nil
}

// RemoveSlide removes the slide into the database
// and returns true, nil if the value was removed; otherwise, false and the error
func (r *repo) RemoveSlide(ctx context.Context, slideID uint64) (bool, error) {
	query := squirrel.Delete(tableName).
		Where(squirrel.Eq{"id": slideID}).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	result, err := query.ExecContext(ctx)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if rowsAffected <= 0 {
		return false, ErrSlideNotFound
	}

	return true, err
}
