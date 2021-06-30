package flusher_test

import (
	"context"
	"errors"

	"github.com/ozoncp/ocp-presentation-api/internal/ocp-slide-api/flusher"
	"github.com/ozoncp/ocp-presentation-api/internal/ocp-slide-api/mock"
	"github.com/ozoncp/ocp-presentation-api/internal/ocp-slide-api/model"

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

		slides []model.Slide
		rest   []model.Slide

		ctx       context.Context
		f         flusher.Flusher
		chunkSize uint

		testCases = []model.Slide{
			{ID: 1, PresentationID: 1},
			{ID: 2, PresentationID: 2},
			{ID: 3, PresentationID: 3},
			{ID: 4, PresentationID: 4},
			{ID: 5, PresentationID: 5},
			{ID: 6, PresentationID: 6},
			{ID: 7, PresentationID: 7},
			{ID: 8, PresentationID: 8},
			{ID: 9, PresentationID: 9},
			{ID: 10, PresentationID: 10},
		}
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())

		mockRepo = mock.NewMockRepo(ctrl)
		ctx = context.Background()
	})

	JustBeforeEach(func() {
		f = flusher.NewFlusher(chunkSize, mockRepo)
		rest, err = f.Flush(ctx, slides)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	// ////////////////////////////////////////////////////////////////////////

	Context("when the flusher can flush all slides", func() {
		BeforeEach(func() {
			slides = make([]model.Slide, len(testCases))
			copy(slides, testCases)

			chunkSize = uint(len(slides) + 1)

			number := int64(len(slides))
			mockRepo.EXPECT().MultiCreateSlides(ctx, gomock.Any()).Return(number, nil)
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

	Context("when flusher can flush a part of the slides", func() {
		var halfSize int64

		BeforeEach(func() {
			slides = make([]model.Slide, len(testCases))
			copy(slides, testCases)

			halfSize = int64(len(slides) / 2)
			chunkSize = uint(halfSize)

			gomock.InOrder(
				mockRepo.EXPECT().MultiCreateSlides(ctx, gomock.Len(int(chunkSize))).Return(halfSize, nil),
				mockRepo.EXPECT().MultiCreateSlides(ctx, gomock.Any()).Return(int64(0), errUnsupportedOperation),
			)
		})

		for i := 0; i < n; i++ {
			It("returns a half slides ", func() {
				Expect(rest).Should(BeEquivalentTo(slides[halfSize:]))
			})

			It("should be the error", func() {
				Expect(err).Should(MatchError(errUnsupportedOperation))
			})
		}
	})

	// ////////////////////////////////////////////////////////////////////////

	Context("when the flusher cannot flush the slides", func() {
		BeforeEach(func() {
			slides = make([]model.Slide, len(testCases))
			copy(slides, testCases)

			chunkSize = 2

			mockRepo.EXPECT().MultiCreateSlides(ctx, gomock.Len(int(chunkSize))).Return(int64(0), errUnsupportedOperation)
		})

		for i := 0; i < n; i++ {
			It("should be equivalent to the slides", func() {
				Expect(rest).Should(BeEquivalentTo(slides))
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
