// Package model represents model data for Ozon Code Platform Presentation API.
package model

import "fmt"

// Presentation represents the connection between a lesson and a slide
type Presentation struct {
	ID          uint64 `db:"id"`
	LessonID    uint64 `db:"lesson_id"`
	UserID      uint64 `db:"user_id"`
	Name        string `db:"name"`
	Description string `db:"description"`
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
