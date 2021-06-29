// Package saver flushes presentations by the time.
package saver

import (
	"context"
	"time"

	"github.com/ozoncp/ocp-presentation-api/internal/common/alarm"
	"github.com/ozoncp/ocp-presentation-api/internal/ocp-presentation-api/flusher"
	"github.com/ozoncp/ocp-presentation-api/internal/ocp-presentation-api/model"
)

const duration = 50 * time.Millisecond

type Saver interface {
	Init(ctx context.Context)
	Save(ctx context.Context, presentation model.Presentation) error
	Close()
}

type saver struct {
	presentation chan model.Presentation
	alarm        alarm.Alarm
	flusher      flusher.Flusher
	done         chan struct{}
}

func NewSaver(capacity int, alarm alarm.Alarm, flusher flusher.Flusher) Saver {
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
			return nil
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
			time.Sleep(duration)
		}
	}
}

func (s *saver) Close() {
	<-s.done
}
