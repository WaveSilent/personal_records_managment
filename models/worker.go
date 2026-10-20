package model

import "uuid"

type Worker struct {
	ID       uuid.UUID
	FullName string
	Position string
}

func NewWorker(fullName string, position string) Worker {
	return Worker{
		ID:       uuid.New(),
		FullName: fullName,
		Position: position,
	}
}
