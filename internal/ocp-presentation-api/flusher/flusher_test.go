package flusher_test

import (
	"context"
	"errors"

	"github.com/ozoncp/ocp-presentation-api/internal/ocp-presentation-api/flusher"
	"github.com/ozoncp/ocp-presentation-api/internal/ocp-presentation-api/mock"
	"github.com/ozoncp/ocp-presentation-api/internal/ocp-presentation-api/model"

	"github.com/golang/mock/gomock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var errUnsupportedOperation = errors.New("not implemented yet")

var n = 10
var _ = Describe("Flusher", func() {

	var (
		err error

		ctrl     *gomock.Controller
		mockRepo *mock.MockRepo

		presentations []model.Presentation
		rest          []model.Presentation

		ctx       context.Context
		f         flusher.Flusher
		chunkSize uint

		testCases = []model.Presentation{
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
		ctrl = gomock.NewController(GinkgoT())

		mockRepo = mock.NewMockRepo(ctrl)
		ctx = context.Background()
	})

	JustBeforeEach(func() {
		f = flusher.NewFlusher(chunkSize, mockRepo)
		rest, err = f.Flush(ctx, presentations)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	// ////////////////////////////////////////////////////////////////////////

	Context("when the flusher can flush all presentations", func() {
		BeforeEach(func() {
			presentations = make([]model.Presentation, len(testCases))
			copy(presentations, testCases)

			chunkSize = uint(len(presentations) + 1)

			number := int64(len(presentations))
			mockRepo.EXPECT().MultiCreatePresentations(ctx, gomock.Any()).Return(number, nil)
		})

		for i := 0; i < n; i++ {
			It("returns the empty slice", func() {
				Expect(rest).Should(BeNil())
			})

			It("should not be an error", func() {
				Expect(err).Should(BeNil())
			})
		}
	})

	Context("when flusher can flush a part of the presentations", func() {
		var halfSize int64

		BeforeEach(func() {
			presentations = make([]model.Presentation, len(testCases))
			copy(presentations, testCases)

			halfSize = int64(len(presentations) / 2)
			chunkSize = uint(halfSize)

			gomock.InOrder(
				mockRepo.EXPECT().MultiCreatePresentations(ctx, gomock.Len(int(chunkSize))).Return(halfSize, nil),
				mockRepo.EXPECT().MultiCreatePresentations(ctx, gomock.Any()).Return(int64(0), errUnsupportedOperation),
			)
		})

		for i := 0; i < n; i++ {
			It("returns a half slides ", func() {
				Expect(rest).Should(BeEquivalentTo(presentations[halfSize:]))
			})

			It("should be the error", func() {
				Expect(err).Should(MatchError(errUnsupportedOperation))
			})
		}
	})

	// ////////////////////////////////////////////////////////////////////////

	Context("when the flusher cannot flush the slides", func() {
		BeforeEach(func() {
			presentations = make([]model.Presentation, len(testCases))
			copy(presentations, testCases)

			chunkSize = 2

			mockRepo.EXPECT().MultiCreatePresentations(ctx, gomock.Len(int(chunkSize))).Return(int64(0), errUnsupportedOperation)
		})

		for i := 0; i < n; i++ {
			It("should be equivalent to the slides", func() {
				Expect(rest).Should(BeEquivalentTo(presentations))
			})

			It("should be the error", func() {
				Expect(err).Should(MatchError(errUnsupportedOperation))
			})
		}
	})

	Context("when the flusher gives incorrect the parameter", func() {
		BeforeEach(func() {
			chunkSize = 0
		})

		for i := 0; i < n; i++ {
			It("should be nil", func() {
				Expect(rest).Should(BeNil())
			})

			It("should be the error", func() {
				Expect(err).Should(MatchError(flusher.ErrInvalidArgument))
			})
		}
	})
})
