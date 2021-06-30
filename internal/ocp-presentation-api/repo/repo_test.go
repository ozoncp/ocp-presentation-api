package repo_test

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"

	"github.com/ozoncp/ocp-presentation-api/internal/ocp-presentation-api/model"
	"github.com/ozoncp/ocp-presentation-api/internal/ocp-presentation-api/repo"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var errDatabaseConnection = errors.New("error establishing a database connection")

var _ = Describe("Repo", func() {
	const (
		n              = 10
		numberOfFields = 4
	)

	var (
		err error

		mock sqlmock.Sqlmock

		db     *sql.DB
		sqlxDB *sqlx.DB

		ctx        context.Context
		repository repo.Repo

		presentations = []model.Presentation{
			{ID: 1, LessonID: 1, UserID: 1},
			{ID: 2, LessonID: 2, UserID: 2},
			{ID: 3, LessonID: 3, UserID: 3},
			{ID: 4, LessonID: 4, UserID: 4},
			{ID: 5, LessonID: 5, UserID: 5},
			{ID: 6, LessonID: 6, UserID: 6},
			{ID: 7, LessonID: 7, UserID: 7},
			{ID: 8, LessonID: 8, UserID: 8},
			{ID: 9, LessonID: 9, UserID: 9},
			{ID: 10, LessonID: 10, UserID: 10},
		}
	)

	BeforeEach(func() {
		db, mock, err = sqlmock.New()
		Expect(err).Should(BeNil())
		sqlxDB = sqlx.NewDb(db, "sqlmock")

		ctx = context.Background()
		repository = repo.NewRepo(sqlxDB)
	})

	JustBeforeEach(func() {
	})

	AfterEach(func() {
		mock.ExpectClose()
		err = db.Close()
		Expect(err).Should(BeNil())
	})

	// ////////////////////////////////////////////////////////////////////////

	Context("when the repository saves the new presentation successfully", func() {
		for i, presentation := range presentations {
			mockID := uint64(i)
			presentation := presentation
			var id uint64

			BeforeEach(func() {
				rows := sqlmock.NewRows([]string{
					"id",
				})
				rows.AddRow(
					mockID,
				)

				query := mock.ExpectQuery("INSERT INTO presentation")
				query.WithArgs(
					presentation.LessonID,
					presentation.UserID,
					presentation.Name,
					presentation.Description,
				)
				query.WillReturnRows(rows)

				id, err = repository.CreatePresentation(ctx, presentation)
			})

			It("should return the ID correctly", func() {
				Expect(id).To(Equal(mockID))
			})

			It("should not be an error", func() {
				Expect(err).Should(BeNil())
			})
		}
	})

	Context("when the repository fails to save the new presentation", func() {
		for _, presentation := range presentations {
			presentation := presentation
			var id uint64

			BeforeEach(func() {
				query := mock.ExpectQuery("INSERT INTO presentation")
				query.WithArgs(
					presentation.LessonID,
					presentation.UserID,
					presentation.Name,
					presentation.Description,
				)
				query.WillReturnError(errDatabaseConnection)

				id, err = repository.CreatePresentation(ctx, presentation)
			})

			It("should return the zero-value for the ID", func() {
				Expect(id).To(BeZero())
			})

			It("should be the error", func() {
				Expect(err).Should(MatchError(errDatabaseConnection))
			})
		}
	})

	Context("when the repository saves the new presentations successfully", func() {
		for i := 0; i < n; i++ {
			var (
				lastInsertID                    int64
				rowsAffected                    int64
				numberOfTheCreatedPresentations int64
			)

			BeforeEach(func() {
				rowsAffected = int64(len(presentations))

				values := make([]driver.Value, numberOfFields*rowsAffected)
				for i, presentation := range presentations {
					lastInsertID = int64(presentation.ID)

					values[numberOfFields*i] = presentation.LessonID
					values[numberOfFields*i+1] = presentation.UserID
					values[numberOfFields*i+2] = presentation.Name
					values[numberOfFields*i+3] = presentation.Description
				}

				mock.ExpectExec("INSERT INTO presentation").
					WithArgs(values...).
					WillReturnResult(sqlmock.NewResult(lastInsertID, rowsAffected))

				numberOfTheCreatedPresentations, err = repository.MultiCreatePresentations(ctx, presentations)
			})

			It("should return the number of the created presentations correctly", func() {
				Expect(numberOfTheCreatedPresentations).To(Equal(rowsAffected))
			})

			It("should not be an error", func() {
				Expect(err).Should(BeNil())
			})
		}
	})

	Context("when the repository fails to save the new presentations", func() {
		for i := 0; i < n; i++ {
			var (
				rowsAffected                    int64
				numberOfTheCreatedPresentations int64
			)

			BeforeEach(func() {
				rowsAffected = int64(len(presentations))

				values := make([]driver.Value, numberOfFields*rowsAffected)
				for i, presentation := range presentations {
					values[numberOfFields*i] = presentation.LessonID
					values[numberOfFields*i+1] = presentation.UserID
					values[numberOfFields*i+2] = presentation.Name
					values[numberOfFields*i+3] = presentation.Description
				}

				mock.ExpectExec("INSERT INTO presentation").WithArgs(values...).WillReturnError(errDatabaseConnection)

				numberOfTheCreatedPresentations, err = repository.MultiCreatePresentations(ctx, presentations)
			})

			It("should return the zero-value for the number of the created presentations", func() {
				Expect(numberOfTheCreatedPresentations).To(BeZero())
			})

			It("should be the error", func() {
				Expect(err).Should(MatchError(errDatabaseConnection))
			})
		}
	})

	// ////////////////////////////////////////////////////////////////////////

	Context("when the repository updates the presentation successfully", func() {
		for _, presentation := range presentations {
			presentation := presentation
			var found bool

			BeforeEach(func() {
				query := mock.ExpectExec("UPDATE presentation SET")
				query.WithArgs(
					presentation.LessonID,
					presentation.UserID,
					presentation.Name,
					presentation.Description,
					presentation.ID,
				)
				query.WillReturnResult(sqlmock.NewResult(1, 1))

				found, err = repository.UpdatePresentation(ctx, presentation)
			})

			It("should return true", func() {
				Expect(found).Should(BeTrue())
			})

			It("should not be an error", func() {
				Expect(err).Should(BeNil())
			})
		}
	})

	Context("when the repository fails to update the presentation", func() {
		for _, presentation := range presentations {
			presentation := presentation
			var found bool

			BeforeEach(func() {
				query := mock.ExpectExec("UPDATE presentation SET")
				query.WithArgs(
					presentation.LessonID,
					presentation.UserID,
					presentation.Name,
					presentation.Description,
					presentation.ID,
				)
				query.WillReturnError(errDatabaseConnection)

				found, err = repository.UpdatePresentation(ctx, presentation)
			})

			It("should return false", func() {
				Expect(found).Should(BeFalse())
			})

			It("should be the error", func() {
				Expect(err).Should(MatchError(errDatabaseConnection))
			})
		}
	})

	Context("when the repository updates the presentation unsuccessfully", func() {
		for _, presentation := range presentations {
			presentation := presentation
			var found bool

			BeforeEach(func() {
				query := mock.ExpectExec("UPDATE presentation SET")
				query.WithArgs(
					presentation.LessonID,
					presentation.UserID,
					presentation.Name,
					presentation.Description,
					presentation.ID,
				)
				query.WillReturnResult(sqlmock.NewResult(0, 0))

				found, err = repository.UpdatePresentation(ctx, presentation)
			})

			It("should return false", func() {
				Expect(found).Should(BeFalse())
			})

			It("should be the error", func() {
				Expect(err).Should(MatchError(repo.ErrPresentationNotFound))
			})
		}
	})

	// ////////////////////////////////////////////////////////////////////////

	Context("when the repository describes the presentation successfully", func() {
		for _, presentation := range presentations {
			presentation := presentation
			var result *model.Presentation

			BeforeEach(func() {
				rows := sqlmock.NewRows([]string{
					"id",
					"lesson_id",
					"user_id",
					"name",
					"description",
				})

				rows.AddRow(
					presentation.ID,
					presentation.LessonID,
					presentation.UserID,
					presentation.Name,
					presentation.Description,
				)

				query := mock.ExpectQuery("SELECT id, lesson_id, user_id, name, description FROM presentation WHERE")
				query.WithArgs(
					presentation.ID,
				)
				query.WillReturnRows(rows)

				result, err = repository.DescribePresentation(ctx, presentation.ID)
			})

			It("should populate the presentation correctly", func() {
				Expect(*result).Should(BeEquivalentTo(presentation))
			})

			It("should not be an error", func() {
				Expect(err).Should(BeNil())
			})
		}
	})

	Context("when the repository fails to describe the presentation", func() {
		for _, presentation := range presentations {
			presentation := presentation
			var result *model.Presentation

			BeforeEach(func() {
				rows := sqlmock.NewRows([]string{
					"id",
					"lesson_id",
					"user_id",
					"name",
					"description",
				})

				rows.AddRow(
					presentation.ID,
					presentation.LessonID,
					presentation.UserID,
					presentation.Name,
					presentation.Description,
				)

				query := mock.ExpectQuery("SELECT id, lesson_id, user_id, name, description FROM presentation WHERE")
				query.WithArgs(
					presentation.ID,
				)
				query.WillReturnError(errDatabaseConnection)

				result, err = repository.DescribePresentation(ctx, presentation.ID)
			})

			It("should be nil", func() {
				Expect(result).Should(BeNil())
			})

			It("should be the error", func() {
				Expect(err).Should(MatchError(errDatabaseConnection))
			})
		}
	})

	Context("when the repository describes the presentation unsuccessfully", func() {
		for _, presentation := range presentations {
			presentation := presentation
			var result *model.Presentation

			BeforeEach(func() {
				query := mock.ExpectQuery("SELECT id, lesson_id, user_id, name, description FROM presentation WHERE")
				query.WithArgs(
					presentation.ID,
				)
				query.WillReturnError(sql.ErrNoRows)

				result, err = repository.DescribePresentation(ctx, presentation.ID)
			})

			It("should be nil", func() {
				Expect(result).Should(BeNil())
			})

			It("should be the error", func() {
				Expect(err).Should(MatchError(repo.ErrPresentationNotFound))
			})
		}
	})

	// ////////////////////////////////////////////////////////////////////////

	Context("when the repository returns the list of presentations successfully", func() {
		const (
			maxLimit = 15
			offset   = 0
		)

		for limit := 1; limit <= maxLimit; limit++ {
			limit := limit
			var result []model.Presentation

			BeforeEach(func() {
				rows := sqlmock.NewRows([]string{
					"id",
					"lesson_id",
					"user_id",
					"name",
					"description",
				})

				for i, presentation := range presentations {
					if i == limit {
						break
					}
					rows.AddRow(
						presentation.ID,
						presentation.LessonID,
						presentation.UserID,
						presentation.Name,
						presentation.Description,
					)
				}

				query := fmt.Sprintf(
					"SELECT id, lesson_id, user_id, name, description FROM presentation LIMIT %d OFFSET %d",
					limit,
					offset)
				mock.ExpectQuery(query).WillReturnRows(rows)

				result, err = repository.ListPresentations(ctx, uint64(limit), offset)
			})

			It("should populate the slice correctly", func() {
				Expect(len(result)).Should(BeEquivalentTo(min(limit, len(result))))
			})

			It("should not be an error", func() {
				Expect(err).Should(BeNil())
			})
		}
	})

	Context("when the repository fails to return the list of presentations", func() {
		const (
			limit  = 10
			offset = 0
		)

		for i := 0; i < n; i++ {
			var result []model.Presentation

			BeforeEach(func() {
				query := fmt.Sprintf(
					"SELECT id, lesson_id, user_id, name, description FROM presentation LIMIT %d OFFSET %d",
					limit,
					offset)
				mock.ExpectQuery(query).WillReturnError(errDatabaseConnection)

				result, err = repository.ListPresentations(ctx, limit, offset)
			})

			It("should return the empty list of the presentations", func() {
				Expect(result).Should(BeNil())
			})

			It("should be the error", func() {
				Expect(err).Should(MatchError(errDatabaseConnection))
			})
		}
	})

	Context("when the repository returns the list of presentations unsuccessfully", func() {
		const (
			limit  = 10
			offset = 0
		)

		for i := 0; i < n; i++ {
			var result []model.Presentation

			BeforeEach(func() {
				query := mock.ExpectQuery(fmt.Sprintf(
					"SELECT id, lesson_id, user_id, name, description FROM presentation LIMIT %d OFFSET %d",
					limit,
					offset))
				query.WillReturnError(sql.ErrNoRows)

				result, err = repository.ListPresentations(ctx, limit, offset)
			})

			It("should return the empty list of presentations", func() {
				Expect(result).Should(BeEmpty())
			})

			It("should be the error", func() {
				Expect(err).Should(MatchError(repo.ErrPresentationNotFound))
			})
		}
	})

	// ////////////////////////////////////////////////////////////////////////

	Context("when the repository removes the presentation successfully", func() {
		for _, presentation := range presentations {
			presentation := presentation
			var found bool

			BeforeEach(func() {
				query := mock.ExpectExec("DELETE FROM presentation WHERE")
				query.WithArgs(
					presentation.ID,
				)
				query.WillReturnResult(sqlmock.NewResult(1, 1))

				found, err = repository.RemovePresentation(ctx, presentation.ID)
			})

			It("should return true", func() {
				Expect(found).Should(BeTrue())
			})

			It("should not be an error", func() {
				Expect(err).Should(BeNil())
			})
		}
	})

	Context("when the repository fails to remove the presentation", func() {
		for _, presentation := range presentations {
			presentation := presentation
			var found bool

			BeforeEach(func() {
				query := mock.ExpectExec("DELETE FROM presentation WHERE")
				query.WithArgs(
					presentation.ID,
				)
				query.WillReturnError(errDatabaseConnection)

				found, err = repository.RemovePresentation(ctx, presentation.ID)
			})

			It("should return false", func() {
				Expect(found).Should(BeFalse())
			})

			It("should be the error", func() {
				Expect(err).Should(MatchError(errDatabaseConnection))
			})
		}
	})

	Context("when the repository removes the presentation unsuccessfully", func() {
		for _, presentation := range presentations {
			presentation := presentation
			var found bool

			BeforeEach(func() {
				query := mock.ExpectExec("DELETE FROM presentation WHERE")
				query.WithArgs(
					presentation.ID,
				)
				query.WillReturnResult(sqlmock.NewResult(0, 0))

				found, err = repository.RemovePresentation(ctx, presentation.ID)
			})

			It("should return false", func() {
				Expect(found).Should(BeFalse())
			})

			It("should be the error", func() {
				Expect(err).Should(MatchError(repo.ErrPresentationNotFound))
			})
		}
	})
})

func min(lhs int, rhs int) int {
	if lhs <= rhs {
		return lhs
	}

	return rhs
}
