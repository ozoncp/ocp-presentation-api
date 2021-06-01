package flusher_test

import (
	"errors"

	flusher "github.com/ozoncp/ocp-presentation-api/internal/flusher/slide"
	"github.com/ozoncp/ocp-presentation-api/internal/mock/slide/mock"
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

		mockRepo *mock.MockRepo

		slides []model.Slide
		rest   []model.Slide

		f flusher.Flusher

		chunkSize uint
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())

		mockRepo = mock.NewMockRepo(ctrl)
	})

	JustBeforeEach(func() {
		f = flusher.NewFlusher(chunkSize, mockRepo)
		rest, err = f.Flush(slides)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("repo saves all slides", func() {

		BeforeEach(func() {
			chunkSize = 2
			slides = []model.Slide{{}}

			mockRepo.EXPECT().AddSlides(gomock.Any()).Return(nil).MinTimes(1)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(rest).Should(BeNil())
		})
	})

	Context("repo does not save slides", func() {

		BeforeEach(func() {
			chunkSize = 2
			slides = []model.Slide{{}, {}}

			mockRepo.EXPECT().AddSlides(gomock.Len(int(chunkSize))).Return(errUnsupportedOperation)
		})

		It("", func() {
			// Expect(err).Should(BeNil())
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
				mockRepo.EXPECT().AddSlides(gomock.Len(int(chunkSize))).Return(nil),
				mockRepo.EXPECT().AddSlides(gomock.Any()).Return(errUnsupportedOperation),
			)
		})

		It("", func() {
			// Expect(err).Should(BeNil())
			Expect(rest).Should(BeEquivalentTo(slides[halfSize:]))
		})
	})
})
