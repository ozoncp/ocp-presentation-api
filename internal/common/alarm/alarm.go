// Package alarm implements a simple internal library for Ozon Code Platform Presentation API.
package alarm

import (
	"context"
	"time"

	"github.com/ozoncp/ocp-presentation-api/internal/common/clockwrapper"
)

type Alarm interface {
	Alarm() <-chan struct{}
	Init()
	Close()
}

func NewAlarm(ctx context.Context, timeout time.Duration, clock clockwrapper.ClockWrapper) Alarm {
	alarms := make(chan struct{})
	done := make(chan struct{})

	return &alarm{
		ctx:     ctx,
		timeout: timeout,
		clock:   clock,
		alarms:  alarms,
		done:    done,
	}
}

type alarm struct {
	ctx     context.Context
	timeout time.Duration
	clock   clockwrapper.ClockWrapper
	alarms  chan struct{}
	done    chan struct{}
}

func (a *alarm) Alarm() <-chan struct{} {
	return a.alarms
}

func (a *alarm) Init() {
	go func() {
		timer := time.After(a.timeout)

		for {
			select {
			case <-timer:
				a.alarms <- struct{}{}
				timer = time.After(a.timeout)
			case <-a.ctx.Done():
				a.done <- struct{}{}
				close(a.alarms)
				close(a.done)
				return
			}
		}
	}()
}

func (a *alarm) Close() {
	<-a.done
}
