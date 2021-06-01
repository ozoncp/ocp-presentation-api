package flusher_test

import (
	"errors"

	flusher "github.com/ozoncp/ocp-presentation-api/internal/flusher/presentation"
	"github.com/ozoncp/ocp-presentation-api/internal/mock/presentation/mock"
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

		presentations []model.Presentation
		rest          []model.Presentation

		f flusher.Flusher

		chunkSize uint
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())

		mockRepo = mock.NewMockRepo(ctrl)
	})

	JustBeforeEach(func() {
		f = flusher.NewFlusher(chunkSize, mockRepo)
		rest, err = f.Flush(presentations)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("repo saves all presentations", func() {

		BeforeEach(func() {
			chunkSize = 2
			presentations = []model.Presentation{{}}

			mockRepo.EXPECT().AddPresentations(gomock.Any()).Return(nil).MinTimes(1)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(rest).Should(BeNil())
		})
	})

	Context("repo does not save presentations", func() {

		BeforeEach(func() {
			chunkSize = 2
			presentations = []model.Presentation{{}, {}}

			mockRepo.EXPECT().AddPresentations(gomock.Len(int(chunkSize))).Return(errUnsupportedOperation)
		})

		It("", func() {
			// Expect(err).Should(BeNil())
			Expect(rest).Should(BeEquivalentTo(presentations))
		})
	})

	Context("repo saves half presentations", func() {

		var (
			halfSize uint
		)

		BeforeEach(func() {
			presentations = []model.Presentation{{}, {}}
			halfSize = uint(len(presentations) / 2)
			chunkSize = halfSize

			gomock.InOrder(
				mockRepo.EXPECT().AddPresentations(gomock.Len(int(chunkSize))).Return(nil),
				mockRepo.EXPECT().AddPresentations(gomock.Any()).Return(errUnsupportedOperation),
			)
		})

		It("", func() {
			// Expect(err).Should(BeNil())
			Expect(rest).Should(BeEquivalentTo(presentations[halfSize:]))
		})
	})
})
