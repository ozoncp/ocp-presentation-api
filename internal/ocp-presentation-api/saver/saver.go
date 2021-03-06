// Package saver flushes presentations by the time.
package saver

import (
	"context"
	"time"

	"github.com/ozoncp/ocp-presentation-api/internal/common/alarm"
	"github.com/ozoncp/ocp-presentation-api/internal/ocp-presentation-api/flusher"
	"github.com/ozoncp/ocp-presentation-api/internal/ocp-presentation-api/model"
)

// Saver is the interface that wraps the methods which save presentations asynchronous.
type Saver interface {
	Init(ctx context.Context)
	Save(ctx context.Context, presentation model.Presentation) error
	Close()
}

type saver struct {
	presentation chan model.Presentation
	duration     time.Duration
	alarm        alarm.Alarm
	flusher      flusher.Flusher
	done         chan struct{}
}

// NewSaver returns the Saver interface
func NewSaver(capacity uint, duration time.Duration, alarm alarm.Alarm, flusher flusher.Flusher) Saver {
	presentation := make(chan model.Presentation, capacity)
	done := make(chan struct{})

	return &saver{
		presentation: presentation,
		alarm:        alarm,
		flusher:      flusher,
		done:         done,
	}
}

func (s *saver) Init(ctx context.Context) {
	go s.flushing(ctx)
}

func (s *saver) Save(ctx context.Context, presentation model.Presentation) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case s.presentation <- presentation:
			return nil
		}
	}
}

func (s *saver) flushing(ctx context.Context) {
	var presentations []model.Presentation

	alarms := s.alarm.Alarm()

	for {
		select {
		case presentation := <-s.presentation:
			presentations = append(presentations, presentation)

		case <-ctx.Done():
			_, _ = s.flusher.Flush(ctx, presentations)
			s.done <- struct{}{}
			close(s.done)
			return

		case <-alarms:
			presentations, _ = s.flusher.Flush(ctx, presentations)

		default:
			time.Sleep(s.duration)
		}
	}
}

func (s *saver) Close() {
	<-s.done
}
