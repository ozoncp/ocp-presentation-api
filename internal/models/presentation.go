package models

// Presentation
type Presentation struct {
	ID          uint64  `json:"id,omitempty"`
	UserID      uint64  `json:"user_id,omitempty"`
	LessonID    uint64  `json:"lesson_id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Slides      []Slide `json:"slides,omitempty"`
}
