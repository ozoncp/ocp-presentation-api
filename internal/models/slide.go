package models

import "fmt"

type SlideType int

const (
	Question SlideType = iota
	Video
	Document
)

// Slide
type Slide struct {
	ID             uint64    `json:"id,omitempty"`
	PresentationID uint64    `json:"presentation_id,omitempty"`
	Number         uint64    `json:"number,omitempty"`
	Type           SlideType `json:"type,omitempty"`
}

func (slideType SlideType) String() string {
	switch slideType {
	case Question:
		return "Question"
	case Video:
		return "Video"
	case Document:
		return "Document"
	}
	panic("invalid value")
}

func (slide *Slide) String() string {
	return fmt.Sprintf("[%04d] Slide = { Type = %v }", slide.ID, slide.Type)
}
