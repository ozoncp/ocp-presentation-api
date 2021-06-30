// Package clockwrapper implements a simple internal library for Ozon Code Platform Presentation API.
package clockwrapper

import "time"

type ClockWrapper interface {
	Now() time.Time
}

type realClock struct{}

func (realClock) Now() time.Time {
	return time.Now()
}

func NewRealClock() ClockWrapper {
	return &realClock{}
}
