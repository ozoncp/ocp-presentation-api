package flusher_test

import (
	"context"
	"errors"

	flusher "github.com/ozoncp/ocp-presentation-api/internal/flusher/slide"
	mock "github.com/ozoncp/ocp-presentation-api/internal/mock/slide"
	"github.com/ozoncp/ocp-presentation-api/internal/model"

	"github.com/golang/mock/gomock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	errUnsupportedOperation = errors.New("not implemented yet")
)

var _ = Describe("Flusher", func() {

	var (
		err error

		ctrl *gomock.Controller
		ctx  context.Context

		mockRepo *mock.MockRepo

		slides []model.Slide
		rest   []model.Slide

		f flusher.Flusher

		chunkSize uint
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		ctx = context.Background()

		mockRepo = mock.NewMockRepo(ctrl)
	})

	JustBeforeEach(func() {
		f = flusher.NewFlusher(chunkSize, mockRepo)
		rest, err = f.Flush(ctx, slides)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("The parameter chunkSize is incorrect", func() {
		BeforeEach(func() {
			chunkSize = 0
			slides = []model.Slide{{}, {}}
		})

		It("should incorrectly identify and return the error", func() {
			Expect(err).Should(MatchError(flusher.ErrInvalidArgument))
			Expect(rest).Should(BeNil())
		})
	})

	Context("repo saves all slides", func() {
		BeforeEach(func() {
			chunkSize = 2
			slides = []model.Slide{{}}

			mockRepo.EXPECT().AddSlides(ctx, gomock.Any()).Return(nil).MinTimes(1)
		})

		It("should correctly identify and save all slides", func() {
			Expect(err).Should(BeNil())
			Expect(rest).Should(BeNil())
		})
	})

	Context("repo does not save slides", func() {
		BeforeEach(func() {
			chunkSize = 2
			slides = []model.Slide{{}, {}}

			mockRepo.EXPECT().AddSlides(ctx, gomock.Len(int(chunkSize))).Return(errUnsupportedOperation)
		})

		It("should correctly identify and save the part of slides", func() {
			Expect(err).Should(MatchError(errUnsupportedOperation))
			Expect(rest).Should(BeEquivalentTo(slides))
		})
	})

	Context("repo saves half slides", func() {

		var (
			halfSize uint
		)

		BeforeEach(func() {
			slides = []model.Slide{{}, {}}
			halfSize = uint(len(slides) / 2)
			chunkSize = halfSize

			gomock.InOrder(
				mockRepo.EXPECT().AddSlides(ctx, gomock.Len(int(chunkSize))).Return(nil),
				mockRepo.EXPECT().AddSlides(ctx, gomock.Any()).Return(errUnsupportedOperation),
			)
		})

		It("should correctly identify and save the part of slides", func() {
			Expect(err).Should(MatchError(errUnsupportedOperation))
			Expect(rest).Should(BeEquivalentTo(slides[halfSize:]))
		})
	})
})
