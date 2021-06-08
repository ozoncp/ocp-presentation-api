package flusher_test

import (
	"context"
	"errors"

	flusher "github.com/ozoncp/ocp-presentation-api/internal/flusher/presentation"
	mock "github.com/ozoncp/ocp-presentation-api/internal/mock/presentation"
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

		presentations []model.Presentation
		rest          []model.Presentation

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
		rest, err = f.Flush(ctx, presentations)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("The parameter chunkSize is incorrect", func() {
		BeforeEach(func() {
			chunkSize = 0
			presentations = []model.Presentation{{}, {}}
		})

		It("should incorrectly identify and return the error", func() {
			Expect(err).Should(MatchError(flusher.ErrInvalidArgument))
			Expect(rest).Should(BeNil())
		})
	})

	Context("repo saves all presentations", func() {
		BeforeEach(func() {
			chunkSize = 2
			presentations = []model.Presentation{{}}

			mockRepo.EXPECT().AddPresentations(ctx, gomock.Any()).Return(nil).MinTimes(1)
		})

		It("should correctly identify and save all presentations", func() {
			Expect(err).Should(BeNil())
			Expect(rest).Should(BeNil())
		})
	})

	Context("repo does not save presentations", func() {
		BeforeEach(func() {
			chunkSize = 2
			presentations = []model.Presentation{{}, {}}

			mockRepo.EXPECT().AddPresentations(ctx, gomock.Len(int(chunkSize))).Return(errUnsupportedOperation)
		})

		It("should correctly identify and save the part of presentations", func() {
			Expect(err).Should(MatchError(errUnsupportedOperation))
			Expect(rest).Should(BeEquivalentTo(presentations))
		})
	})

	Context("repo saves half presentations", func() {
		var halfSize uint

		BeforeEach(func() {
			presentations = []model.Presentation{{}, {}}
			halfSize = uint(len(presentations) / 2)
			chunkSize = halfSize

			gomock.InOrder(
				mockRepo.EXPECT().AddPresentations(ctx, gomock.Len(int(chunkSize))).Return(nil),
				mockRepo.EXPECT().AddPresentations(ctx, gomock.Any()).Return(errUnsupportedOperation),
			)
		})

		It("should correctly identify and save the part of presentations", func() {
			Expect(err).Should(MatchError(errUnsupportedOperation))
			Expect(rest).Should(BeEquivalentTo(presentations[halfSize:]))
		})
	})
})
