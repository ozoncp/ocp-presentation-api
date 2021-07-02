// Package saver flushes slides by the time.
package saver

import (
	"context"
	"time"

	"github.com/ozoncp/ocp-presentation-api/internal/common/alarm"
	"github.com/ozoncp/ocp-presentation-api/internal/ocp-slide-api/flusher"
	"github.com/ozoncp/ocp-presentation-api/internal/ocp-slide-api/model"
)

// Saver is the interface that wraps the methods which save slides asynchronous.
type Saver interface {
	Init(ctx context.Context)
	Save(ctx context.Context, presentation model.Slide) error
	Close()
}

type saver struct {
	slide    chan model.Slide
	duration time.Duration
	alarm    alarm.Alarm
	flusher  flusher.Flusher
	done     chan struct{}
}

// NewSaver returns the Saver interface
func NewSaver(capacity uint, duration time.Duration, alarm alarm.Alarm, flusher flusher.Flusher) Saver {
	presentation := make(chan model.Slide, capacity)
	done := make(chan struct{})

	return &saver{
		slide:   presentation,
		alarm:   alarm,
		flusher: flusher,
		done:    done,
	}
}

func (s *saver) Init(ctx context.Context) {
	go s.flushing(ctx)
}

func (s *saver) Save(ctx context.Context, slide model.Slide) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case s.slide <- slide:
			return nil
		}
	}
}

func (s *saver) flushing(ctx context.Context) {
	var slides []model.Slide

	alarms := s.alarm.Alarm()

	for {
		select {
		case slide := <-s.slide:
			slides = append(slides, slide)

		case <-ctx.Done():
			_, _ = s.flusher.Flush(ctx, slides)
			s.done <- struct{}{}
			close(s.done)
			return

		case <-alarms:
			slides, _ = s.flusher.Flush(ctx, slides)

		default:
			time.Sleep(s.duration)
		}
	}
}

func (s *saver) Close() {
	<-s.done
}
