package models

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
