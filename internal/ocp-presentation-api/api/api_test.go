package api_test

import (
	"context"
	"database/sql"
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

var _ = Describe("Presentation API Server", func() {
	var (
		err error

		ctx context.Context

		db     *sql.DB
		sqlxDB *sqlx.DB
		mock   sqlmock.Sqlmock

		presentations = []model.Presentation{
			{ID: 1, Name: "1"},
			{ID: 2, Name: "2"},
			{ID: 3, Name: "3"},
			{ID: 4, Name: "4"},
			{ID: 5, Name: "5"},
			{ID: 6, Name: "6"},
			{ID: 7, Name: "7"},
			{ID: 8, Name: "8"},
			{ID: 9, Name: "9"},
			{ID: 10, Name: "10"},
		}

		repository repo.Repo
		server     desc.PresentationAPIServer

		createRequest  *desc.CreatePresentationV1Request
		createResponse *desc.CreatePresentationV1Response

		describeRequest  *desc.DescribePresentationV1Request
		describeResponse *desc.DescribePresentationV1Response

		listRequest  *desc.ListPresentationsV1Request
		listResponse *desc.ListPresentationsV1Response

		removeRequest  *desc.RemovePresentationV1Request
		removeResponse *desc.RemovePresentationV1Response
	)

	BeforeEach(func() {
		ctx = context.Background()

		db, mock, err = sqlmock.New()
		Expect(err).Should(BeNil())
		sqlxDB = sqlx.NewDb(db, "sqlmock")

		repository = repo.NewRepo(sqlxDB)
		server = api.NewPresentationAPI(repository, 2)
	})

	JustBeforeEach(func() {
	})

	AfterEach(func() {
		mock.ExpectClose()
		err = db.Close()
		Expect(err).Should(BeNil())
	})

	// CreatePresentationV1

	Context("creating a presentation", func() {
		var id uint64 = 1

		BeforeEach(func() {
			createRequest = &desc.CreatePresentationV1Request{
				Presentation: &desc.NewPresentation{
					UserId:      1,
					LessonId:    1,
					Name:        "Presentation",
					Description: "Description",
				},
			}

			rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
			mock.ExpectQuery("INSERT INTO presentation").WithArgs(
				createRequest.Presentation.UserId,
				createRequest.Presentation.LessonId,
				createRequest.Presentation.Name,
				createRequest.Presentation.Description).
				WillReturnRows(rows)

			createResponse, err = server.CreatePresentationV1(ctx, createRequest)
		})

		It("should work well", func() {
			Expect(err).Should(BeNil())
			Expect(createResponse.PresentationId).Should(Equal(id))
		})
	})

	Context("creating a presentation with an invalid argument", func() {
		BeforeEach(func() {
			createRequest = &desc.CreatePresentationV1Request{}
			createResponse, err = server.CreatePresentationV1(ctx, createRequest)
		})

		It("should be an error", func() {
			Expect(createResponse).Should(BeNil())
			Expect(status.Convert(err).Code()).Should(Equal(codes.InvalidArgument))
		})
	})

	Context("no database connecting", func() {
		BeforeEach(func() {
			createRequest = &desc.CreatePresentationV1Request{
				Presentation: &desc.NewPresentation{
					UserId:      1,
					LessonId:    1,
					Name:        "Presentation",
					Description: "Description",
				},
			}

			mock.ExpectQuery("INSERT INTO presentation").WithArgs(
				createRequest.Presentation.UserId,
				createRequest.Presentation.LessonId,
				createRequest.Presentation.Name,
				createRequest.Presentation.Description).
				WillReturnError(errors.New("bad connection"))

			createResponse, err = server.CreatePresentationV1(ctx, createRequest)
		})

		It("should be an error", func() {
			Expect(createResponse).Should(BeNil())
			Expect(status.Convert(err).Code()).Should(Equal(codes.ResourceExhausted))
		})
	})

	// ////////////////////////////////////////////////////////////////////////

	Context("describing a presentation", func() {
		presentation := model.Presentation{
			ID:          1,
			LessonID:    2,
			UserID:      3,
			Name:        "Name",
			Description: "Description",
		}

		BeforeEach(func() {
			describeRequest = &desc.DescribePresentationV1Request{
				PresentationId: presentation.ID,
			}

			rows := sqlmock.NewRows([]string{
				"id", "lesson_id", "user_id", "name", "description"}).AddRow(
				presentation.ID,
				presentation.LessonID,
				presentation.UserID,
				presentation.Name,
				presentation.Description)

			mock.ExpectQuery("SELECT (.+) FROM presentation WHERE").
				WithArgs(describeRequest.PresentationId).
				WillReturnRows(rows)

			describeResponse, err = server.DescribePresentationV1(ctx, describeRequest)
		})

		It("should work well", func() {
			Expect(presentation.ID).Should(BeEquivalentTo(describeResponse.Presentation.Id))
			Expect(presentation.LessonID).Should(BeEquivalentTo(describeResponse.Presentation.LessonId))
			Expect(presentation.UserID).Should(Equal(describeResponse.Presentation.UserId))
			Expect(presentation.Name).Should(Equal(describeResponse.Presentation.Name))
			Expect(presentation.Description).Should(Equal(describeResponse.Presentation.Description))
			Expect(err).Should(BeNil())
		})
	})

	Context("describing a presentation with an invalid argument", func() {
		BeforeEach(func() {
			describeRequest = &desc.DescribePresentationV1Request{}
			describeResponse, err = server.DescribePresentationV1(ctx, describeRequest)
		})

		It("should be an error", func() {
			Expect(describeResponse).Should(BeNil())
			Expect(status.Convert(err).Code()).Should(Equal(codes.InvalidArgument))
		})
	})

	Context("no database connecting", func() {
		var id uint64 = 1

		BeforeEach(func() {
			describeRequest = &desc.DescribePresentationV1Request{PresentationId: id}

			mock.ExpectQuery("SELECT (.+) FROM presentation WHERE").
				WithArgs(describeRequest.PresentationId).
				WillReturnError(errors.New("bad connection"))

			describeResponse, err = server.DescribePresentationV1(ctx, describeRequest)
		})

		It("should be an error", func() {
			Expect(describeResponse).Should(BeNil())
			Expect(status.Convert(err).Code()).Should(Equal(codes.ResourceExhausted))
		})
	})

	// ////////////////////////////////////////////////////////////////////////

	Context("the list of presentations", func() {
		var (
			limit  uint64 = 10
			offset uint64 = 0
		)

		BeforeEach(func() {
			listRequest = &desc.ListPresentationsV1Request{
				Limit:  limit,
				Offset: offset,
			}

			rows := sqlmock.NewRows([]string{"id", "lesson_id", "user_id", "name", "description"}).
				AddRow(
					presentations[0].ID,
					presentations[0].LessonID,
					presentations[0].UserID,
					presentations[0].Name,
					presentations[0].Description).
				AddRow(
					presentations[1].ID,
					presentations[1].LessonID,
					presentations[1].UserID,
					presentations[1].Name,
					presentations[1].Description)

			mock.ExpectQuery(
				fmt.Sprintf("SELECT id, lesson_id, user_id, name, description FROM presentation LIMIT %d OFFSET %d",
					limit, offset)).
				WillReturnRows(rows)

			listResponse, err = server.ListPresentationsV1(ctx, listRequest)
		})

		It("should work well", func() {
			Expect(err).Should(BeNil())
			Expect(len(listResponse.Presentations)).Should(Equal(2))
		})
	})

	// ////////////////////////////////////////////////////////////////////////

	Context("the list of presentations with an invalid argument", func() {
		BeforeEach(func() {
			removeRequest = &desc.RemovePresentationV1Request{}
			removeResponse, err = server.RemovePresentationV1(ctx, removeRequest)
		})

		It("should be an error", func() {
			Expect(status.Convert(err).Code()).Should(Equal(codes.InvalidArgument))
		})
	})

	Context("no database connecting", func() {
		var (
			limit  uint64 = 10
			offset uint64 = 0
		)

		BeforeEach(func() {
			listRequest = &desc.ListPresentationsV1Request{
				Limit:  limit,
				Offset: offset,
			}

			mock.ExpectQuery(
				fmt.Sprintf(
					"SELECT id, lesson_id, user_id, name, description FROM presentation LIMIT %d OFFSET %d",
					limit, offset)).
				WillReturnError(errors.New("bad connection"))

			listResponse, err = server.ListPresentationsV1(ctx, listRequest)
		})

		It("should be an error", func() {
			Expect(listResponse).Should(BeNil())
			Expect(status.Convert(err).Code()).Should(Equal(codes.ResourceExhausted))
		})
	})

	// ////////////////////////////////////////////////////////////////////////

	Context("removing a presentation", func() {
		var id uint64 = 1

		BeforeEach(func() {
			removeRequest = &desc.RemovePresentationV1Request{
				PresentationId: id,
			}

			mock.ExpectExec("DELETE FROM presentation").
				WithArgs(removeRequest.PresentationId).WillReturnResult(sqlmock.NewResult(0, 1))

			removeResponse, err = server.RemovePresentationV1(ctx, removeRequest)
		})

		It("should work well", func() {
			Expect(removeResponse.Found).Should(Equal(true))
			Expect(err).Should(BeNil())
		})
	})

	Context("removing a nonexistent presentation", func() {
		var id uint64 = 1

		BeforeEach(func() {
			removeRequest = &desc.RemovePresentationV1Request{PresentationId: id}

			mock.ExpectExec("DELETE FROM presentation").
				WithArgs(removeRequest.PresentationId).WillReturnResult(sqlmock.NewResult(0, 0))

			removeResponse, err = server.RemovePresentationV1(ctx, removeRequest)
		})

		It("should work well", func() {
			Expect(removeResponse.Found).Should(BeFalse())
			Expect(status.Convert(err).Code()).Should(BeEquivalentTo(codes.NotFound))
		})
	})

	Context("removing a presentation with an invalid argument", func() {
		BeforeEach(func() {
			removeRequest = &desc.RemovePresentationV1Request{}
			removeResponse, err = server.RemovePresentationV1(ctx, removeRequest)
		})

		It("should be an error", func() {
			Expect(status.Convert(err).Code()).Should(Equal(codes.InvalidArgument))
		})
	})

	Context("no database connecting", func() {
		var id uint64 = 1

		BeforeEach(func() {
			removeRequest = &desc.RemovePresentationV1Request{PresentationId: id}

			mock.ExpectExec("DELETE FROM presentation").
				WithArgs(removeRequest.PresentationId).WillReturnError(errors.New("bad connection"))

			removeResponse, err = server.RemovePresentationV1(ctx, removeRequest)
		})

		It("should be an error", func() {
			Expect(status.Convert(err).Code()).Should(Equal(codes.ResourceExhausted))
		})
	})
})
