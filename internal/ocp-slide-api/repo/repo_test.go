package repo_test

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"

	"github.com/golang/mock/gomock"
	"github.com/ozoncp/ocp-presentation-api/internal/ocp-slide-api/model"
	"github.com/ozoncp/ocp-presentation-api/internal/ocp-slide-api/repo"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var errDatabaseConnection = errors.New("error establishing a database connection")

var n = 10
var _ = Describe("Repo", func() {
	var (
		err error

		ctrl *gomock.Controller
		mock sqlmock.Sqlmock

		db     *sql.DB
		sqlxDB *sqlx.DB

		ctx        context.Context
		repository repo.Repo

		slides = []model.Slide{
			{ID: 1, PresentationID: 1},
			{ID: 2, PresentationID: 2},
			{ID: 3, PresentationID: 3},
			{ID: 4, PresentationID: 4},
			{ID: 5, PresentationID: 5},
			{ID: 6, PresentationID: 6},
			{ID: 7, PresentationID: 7},
			{ID: 8, PresentationID: 9},
			{ID: 9, PresentationID: 9},
			{ID: 10, PresentationID: 10},
		}
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())

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

		ctrl.Finish()
	})

	// ////////////////////////////////////////////////////////////////////////

	Context("when the repository saves the new slide successfully", func() {
		var (
			mockID uint64
			id     uint64
		)

		for i, slide := range slides {
			mockID = uint64(i)

			BeforeEach(func() {
				rows := sqlmock.NewRows([]string{
					"id",
				})
				rows.AddRow(
					mockID,
				)

				query := mock.ExpectQuery("INSERT INTO slide")
				query.WithArgs(
					slide.PresentationID,
					slide.Number,
					slide.Type,
				)
				query.WillReturnRows(rows)

				id, err = repository.CreateSlide(ctx, slide)
			})

			It("should return the ID correctly", func() {
				Expect(id).To(Equal(mockID))
			})

			It("should not be an error", func() {
				Expect(err).Should(BeNil())
			})
		}
	})

	Context("when the repository fails to save the new slide", func() {
		var id uint64

		for _, slide := range slides {
			BeforeEach(func() {
				query := mock.ExpectQuery("INSERT INTO slide")
				query.WithArgs(
					slide.PresentationID,
					slide.Number,
					slide.Type,
				)
				query.WillReturnError(errDatabaseConnection)

				id, err = repository.CreateSlide(ctx, slide)
			})

			It("should return the zero-value for the ID", func() {
				Expect(id).To(BeZero())
			})

			It("should be the error", func() {
				Expect(err).Should(MatchError(errDatabaseConnection))
			})
		}
	})

	Context("when the repository saves the new slides successfully", func() {
		var (
			lastInsertID             int64
			rowsAffected             int64
			numberOfTheCreatedSlides int64
		)

		BeforeEach(func() {
			rowsAffected = int64(len(slides))

			values := make([]driver.Value, 3*rowsAffected)
			for i, newSlide := range slides {
				lastInsertID = int64(newSlide.ID)

				values[3*i] = newSlide.PresentationID
				values[3*i+1] = newSlide.Number
				values[3*i+2] = newSlide.Type
			}

			mock.ExpectExec("INSERT INTO slide").
				WithArgs(values...).
				WillReturnResult(sqlmock.NewResult(lastInsertID, rowsAffected))

			numberOfTheCreatedSlides, err = repository.MultiCreateSlides(ctx, slides)
		})

		for i := 0; i < n; i++ {
			It("should return the number of the created slides correctly", func() {
				Expect(numberOfTheCreatedSlides).To(Equal(rowsAffected))
			})

			It("should not be an error", func() {
				Expect(err).Should(BeNil())
			})
		}
	})

	Context("when the repository fails to save the new slides", func() {
		var (
			rowsAffected             int64
			numberOfTheCreatedSlides int64
		)

		BeforeEach(func() {
			rowsAffected = int64(len(slides))

			values := make([]driver.Value, 3*rowsAffected)
			for i, newSlide := range slides {
				values[3*i] = newSlide.PresentationID
				values[3*i+1] = newSlide.Number
				values[3*i+2] = newSlide.Type
			}

			mock.ExpectExec("INSERT INTO slide").WithArgs(values...).WillReturnError(errDatabaseConnection)

			numberOfTheCreatedSlides, err = repository.MultiCreateSlides(ctx, slides)
		})

		for i := 0; i < n; i++ {
			It("should return the zero-value for the number of the created slides", func() {
				Expect(numberOfTheCreatedSlides).To(BeZero())
			})

			It("should be the error", func() {
				Expect(err).Should(MatchError(errDatabaseConnection))
			})
		}
	})

	// ////////////////////////////////////////////////////////////////////////

	Context("when the repository updates the slide successfully", func() {
		var found bool

		for _, slide := range slides {
			BeforeEach(func() {
				query := mock.ExpectExec("UPDATE slide SET")
				query.WithArgs(
					slide.PresentationID,
					slide.Number,
					slide.Type,
					slide.ID,
				)
				query.WillReturnResult(sqlmock.NewResult(1, 1))

				found, err = repository.UpdateSlide(ctx, slide)
			})

			for i := 0; i < n; i++ {
				It("should return true", func() {
					Expect(found).Should(BeTrue())
				})

				It("should not be an error", func() {
					Expect(err).Should(BeNil())
				})
			}
		}
	})

	Context("when the repository fails to update the slide", func() {
		var found bool

		for _, slide := range slides {
			BeforeEach(func() {
				query := mock.ExpectExec("UPDATE slide SET")
				query.WithArgs(
					slide.PresentationID,
					slide.Number,
					slide.Type,
					slide.ID,
				)
				query.WillReturnError(errDatabaseConnection)

				found, err = repository.UpdateSlide(ctx, slide)
			})

			It("should return false", func() {
				Expect(found).Should(BeFalse())
			})

			It("should be the error", func() {
				Expect(err).Should(MatchError(errDatabaseConnection))
			})
		}
	})

	Context("when the repository updates the slide unsuccessfully", func() {
		var found bool

		for _, slide := range slides {
			BeforeEach(func() {
				query := mock.ExpectExec("UPDATE slide SET")
				query.WithArgs(
					slide.PresentationID,
					slide.Number,
					slide.Type,
					slide.ID,
				)
				query.WillReturnResult(sqlmock.NewResult(0, 0))

				found, err = repository.UpdateSlide(ctx, slide)
			})

			It("should return false", func() {
				Expect(found).Should(BeFalse())
			})

			It("should be the error", func() {
				Expect(err).Should(MatchError(repo.ErrSlideNotFound))
			})
		}
	})

	// ////////////////////////////////////////////////////////////////////////

	Context("when the repository describes the slide successfully", func() {
		var result *model.Slide

		for _, slide := range slides {
			BeforeEach(func() {
				rows := sqlmock.NewRows([]string{
					"id",
					"presentation_id",
					"number",
					"type",
				})

				rows.AddRow(
					slide.ID,
					slide.PresentationID,
					slide.Number,
					slide.Type,
				)

				query := mock.ExpectQuery("SELECT id, presentation_id, number, type FROM slide WHERE")
				query.WithArgs(
					slide.ID,
				)
				query.WillReturnRows(rows)

				result, err = repository.DescribeSlide(ctx, slide.ID)
			})

			It("should populate the slide correctly", func() {
				Expect(*result).Should(BeEquivalentTo(slide))
			})

			It("should not be an error", func() {
				Expect(err).Should(BeNil())
			})
		}
	})

	Context("when the repository fails to describe the slide", func() {
		var result *model.Slide

		for _, slide := range slides {
			BeforeEach(func() {
				rows := sqlmock.NewRows([]string{
					"id",
					"presentation_id",
					"number",
					"type",
				})

				rows.AddRow(
					slide.ID,
					slide.PresentationID,
					slide.Number,
					slide.Type,
				)

				query := mock.ExpectQuery("SELECT id, presentation_id, number, type FROM slide WHERE")
				query.WithArgs(
					slide.ID,
				)
				query.WillReturnError(errDatabaseConnection)

				result, err = repository.DescribeSlide(ctx, slide.ID)
			})

			It("should be nil", func() {
				Expect(result).Should(BeNil())
			})

			It("should be the error", func() {
				Expect(err).Should(MatchError(errDatabaseConnection))
			})
		}
	})

	Context("when the repository describes the slide unsuccessfully", func() {
		var result *model.Slide

		for _, slide := range slides {
			BeforeEach(func() {
				query := mock.ExpectQuery("SELECT id, presentation_id, number, type FROM slide WHERE")
				query.WithArgs(
					slide.ID,
				)
				query.WillReturnError(sql.ErrNoRows)

				result, err = repository.DescribeSlide(ctx, slide.ID)
			})

			It("should be nil", func() {
				Expect(result).Should(BeNil())
			})

			It("should be the error", func() {
				Expect(err).Should(MatchError(repo.ErrSlideNotFound))
			})
		}
	})

	// ////////////////////////////////////////////////////////////////////////

	Context("when the repository returns the list of slides successfully", func() {
		var (
			maxLimit uint64 = 20
			offset   uint64 = 0

			result []model.Slide
		)

		for limit := uint64(1); limit <= maxLimit; limit++ {
			BeforeEach(func() {
				rows := sqlmock.NewRows([]string{
					"id",
					"presentation_id",
					"number",
					"type",
				})

				for _, slide := range slides {
					rows.AddRow(
						slide.ID,
						slide.PresentationID,
						slide.Number,
						slide.Type,
					)
				}

				query := fmt.Sprintf("SELECT id, presentation_id, number, type FROM slide LIMIT %d OFFSET %d",
					limit, offset)
				mock.ExpectQuery(query).WillReturnRows(rows)

				result, err = repository.ListSlides(ctx, limit, offset)
			})

			It("should populate the slice correctly", func() {
				Expect(uint64(len(result))).Should(BeEquivalentTo(min(limit, uint64(len(result)))))
			})

			It("should not be an error", func() {
				Expect(err).Should(BeNil())
			})
		}
	})

	Context("when the repository fails to return the list of slides", func() {
		var (
			limit  uint64 = 10
			offset uint64 = 0

			result []model.Slide
		)

		BeforeEach(func() {
			rows := sqlmock.NewRows([]string{
				"id",
				"presentation_id",
				"number",
				"type",
			})

			for _, slide := range slides {
				rows.AddRow(
					slide.ID,
					slide.PresentationID,
					slide.Number,
					slide.Type,
				)
			}

			query := fmt.Sprintf("SELECT id, presentation_id, number, type FROM slide LIMIT %d OFFSET %d",
				limit, offset)
			mock.ExpectQuery(query).WillReturnError(errDatabaseConnection)

			result, err = repository.ListSlides(ctx, limit, offset)
		})

		for i := 0; i < n; i++ {
			It("should return the empty list of the slides", func() {
				Expect(result).Should(BeNil())
			})

			It("should be the error", func() {
				Expect(err).Should(MatchError(errDatabaseConnection))
			})
		}
	})

	Context("when the repository returns the list of slides unsuccessfully", func() {
		var (
			limit  uint64 = 10
			offset uint64 = 0

			result []model.Slide
		)

		BeforeEach(func() {
			query := mock.ExpectQuery(
				fmt.Sprintf("SELECT id, presentation_id, number, type FROM slide LIMIT %d OFFSET %d",
					limit, offset))
			query.WillReturnError(sql.ErrNoRows)

			result, err = repository.ListSlides(ctx, limit, offset)
		})

		for i := 0; i < n; i++ {
			It("should return the empty list of slides", func() {
				Expect(result).Should(BeEmpty())
			})

			It("should be the error", func() {
				Expect(err).Should(MatchError(repo.ErrSlideNotFound))
			})
		}
	})

	// ////////////////////////////////////////////////////////////////////////

	Context("when the repository removes the slide successfully", func() {
		var found bool

		for _, slide := range slides {
			BeforeEach(func() {
				query := mock.ExpectExec("DELETE FROM slide WHERE")
				query.WithArgs(
					slide.ID,
				)
				query.WillReturnResult(sqlmock.NewResult(1, 1))

				found, err = repository.RemoveSlide(ctx, slide.ID)
			})

			It("should return true", func() {
				Expect(found).Should(BeTrue())
			})

			It("should not be an error", func() {
				Expect(err).Should(BeNil())
			})
		}
	})

	Context("when the repository fails to remove the slide", func() {
		var found bool

		for _, slide := range slides {
			BeforeEach(func() {
				query := mock.ExpectExec("DELETE FROM slide WHERE")
				query.WithArgs(
					slide.ID,
				)
				query.WillReturnError(errDatabaseConnection)

				found, err = repository.RemoveSlide(ctx, slide.ID)
			})

			It("should return false", func() {
				Expect(found).Should(BeFalse())
			})

			It("should be the error", func() {
				Expect(err).Should(MatchError(errDatabaseConnection))
			})
		}
	})

	Context("when the repository removes the slide unsuccessfully", func() {
		var found bool

		for _, slide := range slides {
			BeforeEach(func() {
				query := mock.ExpectExec("DELETE FROM slide WHERE")
				query.WithArgs(
					slide.ID,
				)
				query.WillReturnResult(sqlmock.NewResult(0, 0))

				found, err = repository.RemoveSlide(ctx, slide.ID)
			})

			It("should return false", func() {
				Expect(found).Should(BeFalse())
			})

			It("should be the error", func() {
				Expect(err).Should(MatchError(repo.ErrSlideNotFound))
			})
		}
	})
})

func min(lhs uint64, rhs uint64) uint64 {
	if lhs <= rhs {
		return lhs
	}

	return rhs
}
