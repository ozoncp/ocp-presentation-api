package models

import "fmt"

// Presentation
type Presentation struct {
	ID          uint64  `json:"id,omitempty"`
	UserID      uint64  `json:"user_id,omitempty"`
	LessonID    uint64  `json:"lesson_id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Slides      []Slide `json:"slides,omitempty"`
}

func (presentation *Presentation) String() string {
	return fmt.Sprintf("[%04d] Presentation %s = { LessonID = %04d\tAuthorID = %04v\tDescription = %s\t Slides = %v }",
		presentation.ID,
		presentation.Name,
		presentation.LessonID,
		presentation.UserID,
		presentation.Description,
		presentation.Slides,
	)
}
