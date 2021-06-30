package api_test

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"

	"github.com/ozoncp/ocp-presentation-api/internal/ocp-presentation-api/api"
	"github.com/ozoncp/ocp-presentation-api/internal/ocp-presentation-api/model"
	"github.com/ozoncp/ocp-presentation-api/internal/ocp-presentation-api/repo"
	desc "github.com/ozoncp/ocp-presentation-api/pkg/ocp-presentation-api"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var errDatabaseConnection = errors.New("error establishing a database connection")

var _ = Describe("Presentation API Server", func() {
	const (
		n              = 10
		numberOfFields = 4
		chunkSize      = 1024
	)

	var (
		db     *sql.DB
		sqlxDB *sqlx.DB
		mock   sqlmock.Sqlmock

		ctx        context.Context
		repository repo.Repo
		server     desc.PresentationAPIServer

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
		var err error

		db, mock, err = sqlmock.New()
		Expect(err).Should(BeNil())
		sqlxDB = sqlx.NewDb(db, "sqlmock")

		ctx = context.Background()
		repository = repo.NewRepo(sqlxDB)
		server = api.NewPresentationAPI(repository, chunkSize)
	})

	JustBeforeEach(func() {
	})

	AfterEach(func() {
		mock.ExpectClose()
		err := db.Close()
		Expect(err).Should(BeNil())
	})

	// ////////////////////////////////////////////////////////////////////////

	Context("when a client creates the new presentation successfully", func() {
		for _, presentation := range presentations {
			presentation := presentation
			var (
				err      error
				request  *desc.CreatePresentationV1Request
				response *desc.CreatePresentationV1Response
			)

			BeforeEach(func() {
				request = &desc.CreatePresentationV1Request{
					Presentation: &desc.NewPresentation{
						LessonId:    presentation.LessonID,
						UserId:      presentation.UserID,
						Name:        presentation.Name,
						Description: presentation.Description,
					},
				}

				rows := sqlmock.NewRows([]string{
					"id",
				})
				rows.AddRow(presentation.ID)

				query := mock.ExpectQuery("INSERT INTO presentation")
				query.WithArgs(
					request.Presentation.LessonId,
					request.Presentation.UserId,
					request.Presentation.Name,
					request.Presentation.Description,
				)
				query.WillReturnRows(rows)

				response, err = server.CreatePresentationV1(ctx, request)
			})

			It("should return an ID correctly", func() {
				Expect(response.PresentationId).To(Equal(presentation.ID))
			})

			It("should not be an error", func() {
				Expect(err).Should(BeNil())
			})
		}
	})

	Context("when a client creates the new presentation unsuccessfully because of an invalid argument", func() {
		for _, presentation := range presentations {
			presentation := presentation
			var (
				err      error
				request  *desc.CreatePresentationV1Request
				response *desc.CreatePresentationV1Response
			)

			BeforeEach(func() {
				request = &desc.CreatePresentationV1Request{
					Presentation: &desc.NewPresentation{
						LessonId:    0 * presentation.LessonID,
						UserId:      0 * presentation.UserID,
						Name:        presentation.Name,
						Description: presentation.Description,
					},
				}
				response, err = server.CreatePresentationV1(ctx, request)
			})

			It("should return an empty response", func() {
				Expect(response).Should(BeNil())
			})

			It("should return an invalid argument error", func() {
				Expect(status.Convert(err).Code()).Should(Equal(codes.InvalidArgument))
			})
		}
	})

	Context("when a client creates the new presentation unsuccessfully because of a database connection error", func() {
		for _, presentation := range presentations {
			presentation := presentation
			var (
				err      error
				request  *desc.CreatePresentationV1Request
				response *desc.CreatePresentationV1Response
			)

			BeforeEach(func() {
				request = &desc.CreatePresentationV1Request{
					Presentation: &desc.NewPresentation{
						LessonId:    presentation.LessonID,
						UserId:      presentation.UserID,
						Name:        presentation.Name,
						Description: presentation.Description,
					},
				}

				query := mock.ExpectQuery("INSERT INTO presentation")
				query.WithArgs(
					request.Presentation.LessonId,
					request.Presentation.UserId,
					request.Presentation.Name,
					request.Presentation.Description,
				)
				query.WillReturnError(errDatabaseConnection)

				response, err = server.CreatePresentationV1(ctx, request)
			})

			It("should return an empty response", func() {
				Expect(response).Should(BeNil())
			})

			It("should return a resource exhausted error", func() {
				Expect(status.Convert(err).Code()).Should(Equal(codes.ResourceExhausted))
			})
		}
	})

	// ////////////////////////////////////////////////////////////////////////

	Context("when a client creates multiple presentations successfully", func() {
		for i := 0; i < n; i++ {
			var (
				err          error
				lastInsertID int64
				rowsAffected int64
				request      *desc.MultiCreatePresentationsV1Request
				response     *desc.MultiCreatePresentationsV1Response
			)

			BeforeEach(func() {
				rowsAffected = int64(len(presentations))
				values := make([]driver.Value, numberOfFields*rowsAffected)
				newPresentations := make([]*desc.NewPresentation, len(presentations))

				for i, presentation := range presentations {
					lastInsertID = int64(i)

					newPresentations[i] = &desc.NewPresentation{
						LessonId:    presentation.LessonID,
						UserId:      presentation.UserID,
						Name:        presentation.Name,
						Description: presentation.Description,
					}

					values[numberOfFields*i] = presentation.LessonID
					values[numberOfFields*i+1] = presentation.UserID
					values[numberOfFields*i+2] = presentation.Name
					values[numberOfFields*i+3] = presentation.Description
				}

				request = &desc.MultiCreatePresentationsV1Request{
					Presentations: newPresentations,
				}

				mock.ExpectExec("INSERT INTO presentation").
					WithArgs(values...).
					WillReturnResult(sqlmock.NewResult(lastInsertID, rowsAffected))

				response, err = server.MultiCreatePresentationsV1(ctx, request)
			})

			It("should return a number of the created presentations correctly", func() {
				Expect(response.NumberOfCreatedPresentations).Should(BeEquivalentTo(len(presentations)))
			})

			It("should not be an error", func() {
				Expect(err).Should(BeNil())
			})
		}
	})

	Context("when a client creates multiple presentations unsuccessfully because of an invalid argument", func() {
		for i := 0; i < n; i++ {
			var (
				err      error
				request  *desc.MultiCreatePresentationsV1Request
				response *desc.MultiCreatePresentationsV1Response
			)

			BeforeEach(func() {
				request = &desc.MultiCreatePresentationsV1Request{
					Presentations: nil,
				}

				response, err = server.MultiCreatePresentationsV1(ctx, request)
			})

			It("should return an empty response", func() {
				Expect(response).Should(BeNil())
			})

			It("should return an invalid argument error", func() {
				Expect(status.Convert(err).Code()).Should(Equal(codes.InvalidArgument))
			})
		}
	})

	Context("when a client creates multiple presentations unsuccessfully because of a database connection error", func() {
		for i := 0; i < n; i++ {
			var (
				err          error
				rowsAffected int64
				request      *desc.MultiCreatePresentationsV1Request
				response     *desc.MultiCreatePresentationsV1Response
			)

			BeforeEach(func() {
				rowsAffected = int64(len(presentations))
				values := make([]driver.Value, numberOfFields*rowsAffected)
				newPresentations := make([]*desc.NewPresentation, len(presentations))

				for i, presentation := range presentations {
					newPresentations[i] = &desc.NewPresentation{
						LessonId:    presentation.LessonID,
						UserId:      presentation.UserID,
						Name:        presentation.Name,
						Description: presentation.Description,
					}

					values[numberOfFields*i] = presentation.LessonID
					values[numberOfFields*i+1] = presentation.UserID
					values[numberOfFields*i+2] = presentation.Name
					values[numberOfFields*i+3] = presentation.Description
				}

				request = &desc.MultiCreatePresentationsV1Request{
					Presentations: newPresentations,
				}

				mock.ExpectExec("INSERT INTO presentation").
					WithArgs(values...).
					WillReturnError(errDatabaseConnection)

				response, err = server.MultiCreatePresentationsV1(ctx, request)
			})

			It("should return an empty response", func() {
				Expect(response).Should(BeNil())
			})

			It("should return a resource exhausted error", func() {
				Expect(status.Convert(err).Code()).Should(Equal(codes.ResourceExhausted))
			})
		}
	})

	// ////////////////////////////////////////////////////////////////////////

	Context("when a client updates the presentation successfully", func() {
		for _, presentation := range presentations {
			presentation := presentation
			var (
				err      error
				request  *desc.UpdatePresentationV1Request
				response *desc.UpdatePresentationV1Response
			)

			BeforeEach(func() {
				request = &desc.UpdatePresentationV1Request{
					Presentation: &desc.Presentation{
						Id:          presentation.ID,
						LessonId:    presentation.LessonID,
						UserId:      presentation.UserID,
						Name:        presentation.Name,
						Description: presentation.Description,
					},
				}

				query := mock.ExpectExec("UPDATE presentation SET")
				query.WithArgs(
					presentation.LessonID,
					presentation.UserID,
					presentation.Name,
					presentation.Description,
					presentation.ID,
				)
				query.WillReturnResult(sqlmock.NewResult(1, 1))

				response, err = server.UpdatePresentationV1(ctx, request)
			})

			It("should return true the presentation correctly", func() {
				Expect(response.Found).Should(BeTrue())
			})

			It("should not be an error", func() {
				Expect(err).Should(BeNil())
			})
		}
	})

	Context("when a client updates the presentation unsuccessfully", func() {
		for _, presentation := range presentations {
			presentation := presentation
			var (
				err      error
				request  *desc.UpdatePresentationV1Request
				response *desc.UpdatePresentationV1Response
			)

			BeforeEach(func() {
				request = &desc.UpdatePresentationV1Request{
					Presentation: &desc.Presentation{
						Id:          presentation.ID,
						LessonId:    presentation.LessonID,
						UserId:      presentation.UserID,
						Name:        presentation.Name,
						Description: presentation.Description,
					},
				}

				query := mock.ExpectExec("UPDATE presentation SET")
				query.WithArgs(
					presentation.LessonID,
					presentation.UserID,
					presentation.Name,
					presentation.Description,
					presentation.ID,
				)
				query.WillReturnResult(sqlmock.NewResult(0, 0))

				response, err = server.UpdatePresentationV1(ctx, request)
			})

			It("should be an empty response", func() {
				Expect(response).To(BeNil())
			})

			It("should return a not found error", func() {
				Expect(status.Convert(err).Code()).Should(Equal(codes.NotFound))
			})
		}
	})

	Context("when a client updates the presentation unsuccessfully because of an invalid argument", func() {
		for i := 0; i < n; i++ {
			var (
				err      error
				request  *desc.UpdatePresentationV1Request
				response *desc.UpdatePresentationV1Response
			)

			BeforeEach(func() {
				request = &desc.UpdatePresentationV1Request{
					Presentation: &desc.Presentation{
						Id:          0,
						LessonId:    0,
						UserId:      0,
						Name:        "",
						Description: "",
					},
				}
				response, err = server.UpdatePresentationV1(ctx, request)
			})

			It("should return an empty response", func() {
				Expect(response).Should(BeNil())
			})

			It("should return an invalid argument error", func() {
				Expect(status.Convert(err).Code()).Should(Equal(codes.InvalidArgument))
			})
		}
	})

	Context("when a client updates the presentation unsuccessfully because of a database connection error", func() {
		for _, presentation := range presentations {
			presentation := presentation
			var (
				err      error
				request  *desc.UpdatePresentationV1Request
				response *desc.UpdatePresentationV1Response
			)

			BeforeEach(func() {
				request = &desc.UpdatePresentationV1Request{
					Presentation: &desc.Presentation{
						Id:          presentation.ID,
						LessonId:    presentation.LessonID,
						UserId:      presentation.UserID,
						Name:        presentation.Name,
						Description: presentation.Description,
					},
				}

				query := mock.ExpectExec("UPDATE presentation SET")
				query.WithArgs(
					presentation.LessonID,
					presentation.UserID,
					presentation.Name,
					presentation.Description,
					presentation.ID,
				)
				query.WillReturnError(errDatabaseConnection)

				response, err = server.UpdatePresentationV1(ctx, request)
			})

			It("should return an empty response", func() {
				Expect(response).Should(BeNil())
			})

			It("should return a resource exhausted error", func() {
				Expect(status.Convert(err).Code()).Should(Equal(codes.ResourceExhausted))
			})
		}
	})

	// ////////////////////////////////////////////////////////////////////////

	Context("when a client gets the presentation successfully", func() {
		for _, presentation := range presentations {
			presentation := presentation
			var (
				err      error
				request  *desc.DescribePresentationV1Request
				response *desc.DescribePresentationV1Response
			)

			BeforeEach(func() {
				request = &desc.DescribePresentationV1Request{
					PresentationId: presentation.ID,
				}

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

				mock.ExpectQuery("SELECT (.+) FROM presentation WHERE").
					WithArgs(request.PresentationId).
					WillReturnRows(rows)

				response, err = server.DescribePresentationV1(ctx, request)
			})

			It("should populate the presentation correctly", func() {
				Expect(response.Presentation.Id).Should(BeEquivalentTo(presentation.ID))
				Expect(response.Presentation.LessonId).Should(BeEquivalentTo(presentation.LessonID))
				Expect(response.Presentation.UserId).Should(Equal(presentation.UserID))
				Expect(response.Presentation.Name).Should(Equal(presentation.Name))
				Expect(response.Presentation.Description).Should(Equal(presentation.Description))
			})

			It("should not be an error", func() {
				Expect(err).Should(BeNil())
			})
		}
	})

	Context("when a client gets the presentation unsuccessfully", func() {
		for _, presentation := range presentations {
			presentation := presentation
			var (
				err      error
				request  *desc.DescribePresentationV1Request
				response *desc.DescribePresentationV1Response
			)

			BeforeEach(func() {
				request = &desc.DescribePresentationV1Request{
					PresentationId: presentation.ID,
				}

				query := mock.ExpectQuery("SELECT id, lesson_id, user_id, name, description FROM presentation WHERE")
				query.WithArgs(
					presentation.ID,
				)
				query.WillReturnError(sql.ErrNoRows)

				response, err = server.DescribePresentationV1(ctx, request)
			})

			It("should be an empty response", func() {
				Expect(response).To(BeNil())
			})

			It("should return a not found error", func() {
				Expect(status.Convert(err).Code()).Should(Equal(codes.NotFound))
			})
		}
	})

	Context("when a client gets the presentation unsuccessfully because of an invalid argument", func() {
		for i := 0; i < n; i++ {
			var (
				err      error
				request  *desc.DescribePresentationV1Request
				response *desc.DescribePresentationV1Response
			)

			BeforeEach(func() {
				request = &desc.DescribePresentationV1Request{
					PresentationId: 0,
				}
				response, err = server.DescribePresentationV1(ctx, request)
			})

			It("should return an empty response", func() {
				Expect(response).Should(BeNil())
			})

			It("should return an invalid argument error", func() {
				Expect(status.Convert(err).Code()).Should(Equal(codes.InvalidArgument))
			})
		}
	})

	Context("when a client gets the presentation unsuccessfully because of a database connection error", func() {
		for _, presentation := range presentations {
			presentation := presentation
			var (
				err      error
				request  *desc.DescribePresentationV1Request
				response *desc.DescribePresentationV1Response
			)

			BeforeEach(func() {
				request = &desc.DescribePresentationV1Request{
					PresentationId: presentation.ID,
				}

				mock.ExpectQuery("SELECT (.+) FROM presentation WHERE").
					WithArgs(request.PresentationId).
					WillReturnError(errDatabaseConnection)

				response, err = server.DescribePresentationV1(ctx, request)
			})

			It("should return an empty response", func() {
				Expect(response).Should(BeNil())
			})

			It("should return a resource exhausted error", func() {
				Expect(status.Convert(err).Code()).Should(Equal(codes.ResourceExhausted))
			})
		}
	})

	// ////////////////////////////////////////////////////////////////////////

	Context("when a client gets the list successfully", func() {
		const (
			maxLimit = 15
			offset   = 0
		)

		for limit := 1; limit <= maxLimit; limit++ {
			limit := limit
			var (
				err      error
				request  *desc.ListPresentationsV1Request
				response *desc.ListPresentationsV1Response
			)

			BeforeEach(func() {
				request = &desc.ListPresentationsV1Request{
					Limit: uint64(limit),
				}

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

				response, err = server.ListPresentationsV1(ctx, request)
			})

			It("should populate the response correctly", func() {
				Expect(len(response.Presentations)).Should(BeEquivalentTo(min(limit, len(presentations))))
			})

			It("should not be an error", func() {
				Expect(err).Should(BeNil())
			})
		}
	})

	Context("when a client gets the list unsuccessfully", func() {
		const (
			limit  = 10
			offset = 0
		)

		for i := 0; i < n; i++ {
			var (
				err      error
				request  *desc.ListPresentationsV1Request
				response *desc.ListPresentationsV1Response
			)

			BeforeEach(func() {
				request = &desc.ListPresentationsV1Request{
					Limit:  limit,
					Offset: offset,
				}

				query := fmt.Sprintf(
					"SELECT id, lesson_id, user_id, name, description FROM presentation LIMIT %d OFFSET %d",
					limit,
					offset)
				mock.ExpectQuery(query).WillReturnError(sql.ErrNoRows)

				response, err = server.ListPresentationsV1(ctx, request)
			})

			It("should be an empty response", func() {
				Expect(response).To(BeNil())
			})

			It("should return a not found error", func() {
				Expect(status.Convert(err).Code()).Should(Equal(codes.NotFound))
			})
		}
	})

	Context("when a client gets the list unsuccessfully because of an invalid argument", func() {
		for i := 0; i < n; i++ {
			var (
				err      error
				request  *desc.ListPresentationsV1Request
				response *desc.ListPresentationsV1Response
			)

			BeforeEach(func() {
				request = &desc.ListPresentationsV1Request{
					Limit:  0,
					Offset: 0,
				}
				response, err = server.ListPresentationsV1(ctx, request)
			})

			It("should be an empty response", func() {
				Expect(response).To(BeNil())
			})

			It("should return an invalid argument error", func() {
				Expect(status.Convert(err).Code()).Should(Equal(codes.InvalidArgument))
			})
		}
	})

	Context("when a client gets the list unsuccessfully because of a database connection error", func() {
		const (
			limit  = 10
			offset = 0
		)

		for i := 0; i < n; i++ {
			var (
				err      error
				request  *desc.ListPresentationsV1Request
				response *desc.ListPresentationsV1Response
			)

			BeforeEach(func() {
				request = &desc.ListPresentationsV1Request{
					Limit:  limit,
					Offset: offset,
				}

				query := mock.ExpectQuery(fmt.Sprintf(
					"SELECT id, lesson_id, user_id, name, description FROM presentation LIMIT %d OFFSET %d",
					limit,
					offset))
				query.WillReturnError(errDatabaseConnection)

				response, err = server.ListPresentationsV1(ctx, request)
			})

			It("should return an empty response", func() {
				Expect(response).Should(BeNil())
			})

			It("should return a resource exhausted error", func() {
				Expect(status.Convert(err).Code()).Should(Equal(codes.ResourceExhausted))
			})
		}
	})

	// ////////////////////////////////////////////////////////////////////////

	Context("when a client removes the presentation successfully", func() {
		for _, presentation := range presentations {
			presentation := presentation
			var (
				err      error
				request  *desc.RemovePresentationV1Request
				response *desc.RemovePresentationV1Response
			)

			BeforeEach(func() {
				request = &desc.RemovePresentationV1Request{
					PresentationId: presentation.ID,
				}

				mock.ExpectExec("DELETE FROM presentation").
					WithArgs(request.PresentationId).
					WillReturnResult(sqlmock.NewResult(0, 1))

				response, err = server.RemovePresentationV1(ctx, request)
			})

			It("should return true", func() {
				Expect(response.Found).To(BeTrue())
			})

			It("should not be an error", func() {
				Expect(err).Should(BeNil())
			})
		}
	})

	Context("when a client removes the presentation unsuccessfully", func() {
		for _, presentation := range presentations {
			presentation := presentation
			var (
				err      error
				request  *desc.RemovePresentationV1Request
				response *desc.RemovePresentationV1Response
			)

			BeforeEach(func() {
				request = &desc.RemovePresentationV1Request{
					PresentationId: presentation.ID,
				}

				mock.ExpectExec("DELETE FROM presentation").
					WithArgs(request.PresentationId).
					WillReturnResult(sqlmock.NewResult(0, 0))

				response, err = server.RemovePresentationV1(ctx, request)
			})

			It("should return false", func() {
				Expect(response.Found).To(BeFalse())
			})

			It("should return a not found error", func() {
				Expect(status.Convert(err).Code()).Should(Equal(codes.NotFound))
			})
		}
	})

	Context("when a client removes the presentation unsuccessfully because of an invalid argument", func() {
		for i := 0; i < n; i++ {
			var (
				err      error
				request  *desc.RemovePresentationV1Request
				response *desc.RemovePresentationV1Response
			)

			BeforeEach(func() {
				request = &desc.RemovePresentationV1Request{
					PresentationId: 0,
				}
				response, err = server.RemovePresentationV1(ctx, request)
			})

			It("should return an empty response", func() {
				Expect(response).Should(BeNil())
			})

			It("should return an invalid argument error", func() {
				Expect(status.Convert(err).Code()).Should(Equal(codes.InvalidArgument))
			})
		}
	})

	Context("when a client removes the presentation unsuccessfully because of a database connection error", func() {
		for _, presentation := range presentations {
			presentation := presentation
			var (
				err      error
				request  *desc.RemovePresentationV1Request
				response *desc.RemovePresentationV1Response
			)

			BeforeEach(func() {
				request = &desc.RemovePresentationV1Request{
					PresentationId: presentation.ID,
				}

				mock.ExpectExec("DELETE FROM presentation").
					WithArgs(request.PresentationId).
					WillReturnError(errDatabaseConnection)

				response, err = server.RemovePresentationV1(ctx, request)
			})

			It("should return an empty response", func() {
				Expect(response).Should(BeNil())
			})

			It("should return a resource exhausted error", func() {
				Expect(status.Convert(err).Code()).Should(Equal(codes.ResourceExhausted))
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
