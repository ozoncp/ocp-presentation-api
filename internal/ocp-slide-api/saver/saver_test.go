package saver_test

import (
	"context"
	"time"

	"github.com/ozoncp/ocp-presentation-api/internal/ocp-slide-api/mock"
	"github.com/ozoncp/ocp-presentation-api/internal/ocp-slide-api/model"
	"github.com/ozoncp/ocp-presentation-api/internal/ocp-slide-api/saver"

	"github.com/golang/mock/gomock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	capacity = 100
	duration = 50 * time.Millisecond
)

var _ = Describe("Saver", func() {
	var (
		err error

		ctrl *gomock.Controller

		mockFlusher *mock.MockFlusher
		mockAlarm   *mock.MockAlarm

		ctx    context.Context
		alarms chan struct{}
		s      saver.Saver

		slide = model.Slide{
			ID:             1,
			PresentationID: 1,
		}
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())

		mockAlarm = mock.NewMockAlarm(ctrl)
		mockFlusher = mock.NewMockFlusher(ctrl)

		alarms = make(chan struct{})
		mockAlarm.EXPECT().Alarm().Return(alarms).AnyTimes()

		ctx = context.Background()
		s = saver.NewSaver(capacity, duration, mockAlarm, mockFlusher)
	})

	JustBeforeEach(func() {
		s.Init(ctx)
		err = s.Save(ctx, slide)
	})

	AfterEach(func() {
		s.Close()
		ctrl.Finish()
	})

	Context("when a context with a deadline can be used", func() {
		var cancelFunc context.CancelFunc

		BeforeEach(func() {
			ctx, cancelFunc = context.WithDeadline(context.Background(), time.Now().Add(time.Millisecond))
			mockFlusher.EXPECT().Flush(gomock.Any(), gomock.Any()).Return(nil, nil).MinTimes(1).MaxTimes(20)
		})

		JustBeforeEach(func() {
			cancelFunc()
		})

		It("should not be an error", func() {
			Expect(err).Should(BeNil())
		})
	})

	Context("when a context can be canceled", func() {
		var cancelFunc context.CancelFunc

		BeforeEach(func() {
			ctx, cancelFunc = context.WithCancel(ctx)
			mockFlusher.EXPECT().Flush(gomock.Any(), gomock.Any()).Return(nil, nil).MinTimes(1).MaxTimes(20)
		})

		JustBeforeEach(func() {
			cancelFunc()
		})

		It("should not be an error", func() {
			Expect(err).Should(BeNil())
		})
	})

	Context("when an alarm can occur", func() {
		var cancelFunc context.CancelFunc

		BeforeEach(func() {
			ctx, cancelFunc = context.WithCancel(ctx)
			mockFlusher.EXPECT().Flush(gomock.Any(), gomock.Any()).Return(nil, nil).MinTimes(1).MaxTimes(20)
		})

		JustBeforeEach(func() {
			alarms <- struct{}{}
			cancelFunc()
		})

		It("should not be an error", func() {
			Expect(err).Should(BeNil())
		})
	})
})
