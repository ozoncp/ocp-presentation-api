// Package models represent model data for Ozon Code Platform Presentation API.
package models

import "fmt"

type ContentType uint8

const (
	TypeDocument ContentType = iota
	TypeVideo
	TypeQuestion
	TypeTask
)

// Slide represents the connection between a presentation and a content
type Slide struct {
	ID             uint64      `json:"id,omitempty"`
	PresentationID uint64      `json:"presentation_id,omitempty"`
	Number         uint64      `json:"number,omitempty"`
	Type           ContentType `json:"type,omitempty"`
}

func (contentType ContentType) String() string {
	switch contentType {
	case TypeDocument:
		return "Document"
	case TypeVideo:
		return "Video"
	case TypeQuestion:
		return "Question"
	case TypeTask:
		return "Task"
	}
	panic("invalid value")
}

func (slide *Slide) String() string {
	return fmt.Sprintf("[%04d] Slide = { Type = %v }", slide.ID, slide.Type)
}
