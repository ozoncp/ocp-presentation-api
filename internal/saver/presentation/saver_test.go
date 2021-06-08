package saver_test

import (
	"context"

	mock "github.com/ozoncp/ocp-presentation-api/internal/mock/presentation"
	"github.com/ozoncp/ocp-presentation-api/internal/model"
	saver "github.com/ozoncp/ocp-presentation-api/internal/saver/presentation"

	"github.com/golang/mock/gomock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	capacity = 100
)

var _ = Describe("Saver", func() {

	var (
		err error

		ctrl *gomock.Controller
		ctx  context.Context

		mockFlusher *mock.MockFlusher
		mockAlarm   *mock.MockAlarm

		presentation model.Presentation
		s            saver.Saver

		alarms chan struct{}
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		ctx = context.Background()

		mockAlarm = mock.NewMockAlarm(ctrl)
		mockFlusher = mock.NewMockFlusher(ctrl)

		alarms = make(chan struct{})
		mockAlarm.EXPECT().Alarm().Return(alarms).AnyTimes()

		s = saver.NewSaver(capacity, mockAlarm, mockFlusher)
	})

	JustBeforeEach(func() {
		s.Init(ctx)
		err = s.Save(ctx, presentation)
	})

	AfterEach(func() {
		s.Close()
		ctrl.Finish()
	})

	Context("ctx canceled", func() {
		var cancelFunc context.CancelFunc

		BeforeEach(func() {
			ctx, cancelFunc = context.WithCancel(ctx)
			mockFlusher.EXPECT().Flush(gomock.Any(), gomock.Any()).Return(nil, nil)
		})

		JustBeforeEach(func() {
			cancelFunc()
		})

		It("should be cancelled", func() {
			Expect(err).Should(BeNil())
		})
	})

	Context("alarm is occurring", func() {
		var cancelFunc context.CancelFunc

		BeforeEach(func() {
			ctx, cancelFunc = context.WithCancel(ctx)
			mockFlusher.EXPECT().Flush(gomock.Any(), gomock.Any()).Return(nil, nil).MinTimes(1).MaxTimes(2)
		})

		JustBeforeEach(func() {
			alarms <- struct{}{}
			cancelFunc()
		})

		It("should flush presentations", func() {
			Expect(err).Should(BeNil())
		})
	})
})
