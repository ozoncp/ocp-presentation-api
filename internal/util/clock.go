// Package util implement a simple internal library for Ozon Code Platform Presentation API.
package util

import "time"

type Clock interface {
	Now() time.Time
}

type realClock struct{}

func (realClock) Now() time.Time {
	return time.Now()
}

func NewRealClock() Clock {
	return &realClock{}
}
