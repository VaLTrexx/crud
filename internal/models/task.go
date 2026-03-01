package models

import "time"

type Task struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Priority int       `json:"priority"`
	Position int       `json:"position"`
	DueDate  time.Time `json:"due_date"`
}
