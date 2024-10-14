package model

import "time"

// Todo is the model for the todo endpoint.
type Todo struct {
	ID        int `gorm:"primaryKey"`
	Task      string
	Status    Status
	Priority  int       `gorm:"default:1;not null;max:10;min:1"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// Status is the status of the task.
type Status string

const (
	// Created is the status for a created task.
	Created = Status("created")
	// Processing is the status for a processing task.
	Processing = Status("processing")
	// Done is the status for a done task.
	Done = Status("done")
)

// StatusMap is a map of task status.
var StatusMap = map[Status]bool{
	Created:    true,
	Processing: true,
	Done:       true,
}

type TodoQuery struct {
	Task      string
	Status    Status
	SortBy    string
	SortOrder string
}
