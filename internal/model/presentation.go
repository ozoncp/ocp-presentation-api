// Package model represents model data for Ozon Code Platform Presentation API.
package model

import "fmt"

// Presentation represents the connection between a lesson and a slide
type Presentation struct {
	ID          uint64 `json:"id,omitempty"`
	LessonID    uint64 `json:"lesson_id,omitempty"`
	UserID      uint64 `json:"user_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

func (presentation *Presentation) String() string {
	return fmt.Sprintf("[%04d] Presentation %s = { LessonID = %04d\tAuthorID = %04v\tDescription = %s\t }",
		presentation.ID,
		presentation.Name,
		presentation.LessonID,
		presentation.UserID,
		presentation.Description,
	)
}
