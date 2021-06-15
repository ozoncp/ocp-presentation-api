// Package repo saves data into the database.
package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-presentation-api/internal/model"

	"github.com/Masterminds/squirrel"
)

const tableName = "presentation"

var (
	ErrPresentationNotFound = errors.New("presentation not found")
	ErrUnknown              = errors.New("unknown")
)

// Repo is the interface that wraps the basic methods of the database.
type Repo interface {
	CreatePresentation(ctx context.Context, presentation model.Presentation) (uint64, error)
	DescribePresentation(ctx context.Context, presentationID uint64) (*model.Presentation, error)
	ListPresentations(ctx context.Context, limit uint64, offset uint64) ([]model.Presentation, error)
	RemovePresentation(ctx context.Context, presentationID uint64) (bool, error)
}

type repo struct {
	db        *sqlx.DB
	chunkSize uint
}

// NewRepo returns the Repo interface
func NewRepo(db *sqlx.DB, chunkSize uint) Repo {
	return &repo{
		db:        db,
		chunkSize: chunkSize,
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

func (r *repo) DescribePresentation(ctx context.Context, presentationID uint64) (*model.Presentation, error) {
	query := squirrel.Select("id", "lesson_id", "user_id", "name", "description").
		From(tableName).
		Where(squirrel.Eq{"id": presentationID}).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	sqlString, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var result []model.Presentation
	err = r.db.SelectContext(ctx, &result, sqlString, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrPresentationNotFound
		}

		return nil, err
	}

	if len(result) == 0 {
		return nil, ErrUnknown
	}

	return &result[0], nil
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
		if err == sql.ErrNoRows {
			return nil, ErrPresentationNotFound
		}

		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	presentations := make([]model.Presentation, 0)

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

	number, err := result.RowsAffected()
	return number != 0, err
}
