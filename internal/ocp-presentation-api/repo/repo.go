// Package repo saves data into the database.
package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ozoncp/ocp-presentation-api/internal/ocp-presentation-api/model"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

const tableName = "presentation"

var (
	ErrPresentationNotFound = errors.New("presentation not found")
	ErrUnknown              = errors.New("unknown")
)

// Repo is the interface that wraps the basic methods of the database.
type Repo interface {
	CreatePresentation(ctx context.Context, presentation model.Presentation) (uint64, error)
	MultiCreatePresentations(ctx context.Context, presentations []model.Presentation) (int64, error)
	UpdatePresentation(ctx context.Context, presentation model.Presentation) (bool, error)
	DescribePresentation(ctx context.Context, presentationID uint64) (*model.Presentation, error)
	ListPresentations(ctx context.Context, limit uint64, offset uint64) ([]model.Presentation, error)
	RemovePresentation(ctx context.Context, presentationID uint64) (bool, error)
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

func (r *repo) CreatePresentation(ctx context.Context, presentation model.Presentation) (uint64, error) {
	query := squirrel.Insert(tableName).
		Columns("lesson_id", "user_id", "name", "description").
		Values(presentation.LessonID, presentation.UserID, presentation.Name, presentation.Description).
		Suffix("RETURNING \"id\"").
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	err := query.QueryRowContext(ctx).Scan(&presentation.ID)
	if err != nil {
		return 0, err
	}

	return presentation.ID, nil
}

func (r *repo) MultiCreatePresentations(ctx context.Context, presentations []model.Presentation) (int64, error) {
	query := squirrel.Insert(tableName).
		Columns("lesson_id", "user_id", "name", "description").
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	for i := range presentations {
		query = query.Values(
			presentations[i].LessonID,
			presentations[i].UserID,
			presentations[i].Name,
			presentations[i].Description,
		)
	}

	result, err := query.ExecContext(ctx)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (r *repo) UpdatePresentation(ctx context.Context, presentation model.Presentation) (bool, error) {
	query := squirrel.Update(tableName).
		Set("lesson_id", presentation.LessonID).
		Set("user_id", presentation.UserID).
		Set("name", presentation.Name).
		Set("description", presentation.Description).
		Where(squirrel.Eq{"id": presentation.ID}).
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
		return false, ErrPresentationNotFound
	}

	return true, nil
}

func (r *repo) DescribePresentation(ctx context.Context, presentationID uint64) (*model.Presentation, error) {
	query := squirrel.Select("id", "lesson_id", "user_id", "name", "description").
		From(tableName).
		Where(squirrel.Eq{"id": presentationID}).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	rows, err := query.QueryContext(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrPresentationNotFound
		}

		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	var presentations []model.Presentation
	for rows.Next() {
		var presentation model.Presentation
		if err = rows.Scan(
			&presentation.ID,
			&presentation.LessonID,
			&presentation.UserID,
			&presentation.Name,
			&presentation.Description); err != nil {
			return nil, err
		}
		presentations = append(presentations, presentation)
	}

	return &presentations[0], nil
}

func (r *repo) ListPresentations(ctx context.Context, limit uint64, offset uint64) ([]model.Presentation, error) {
	query := squirrel.Select("id", "lesson_id", "user_id", "name", "description").
		From(tableName).
		RunWith(r.db).
		Limit(limit).
		Offset(offset).
		PlaceholderFormat(squirrel.Dollar)

	rows, err := query.QueryContext(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrPresentationNotFound
		}

		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	var presentations []model.Presentation
	for rows.Next() {
		var presentation model.Presentation
		if err = rows.Scan(
			&presentation.ID,
			&presentation.LessonID,
			&presentation.UserID,
			&presentation.Name,
			&presentation.Description); err != nil {
			return nil, err
		}
		presentations = append(presentations, presentation)
	}

	return presentations, nil
}

func (r *repo) RemovePresentation(ctx context.Context, presentationID uint64) (bool, error) {
	query := squirrel.Delete(tableName).
		Where(squirrel.Eq{"id": presentationID}).
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
		return false, ErrPresentationNotFound
	}

	return true, err
}
