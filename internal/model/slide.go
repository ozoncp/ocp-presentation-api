// Package model represents model data for Ozon Code Platform Presentation API.
package model

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
	ID             uint64      `db:"id"`
	PresentationID uint64      `db:"presentation_id"`
	Number         uint64      `db:"number"`
	Type           ContentType `db:"type"`
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
