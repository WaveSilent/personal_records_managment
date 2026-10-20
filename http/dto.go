package http

import (
	"encoding/json"
	"time"
	"uuid"
)

type WorkerDTO struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"fullName"`
	Position string    `json:"position"`
}

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func NewErrorDTO(err error) *ErrorDTO {
	return &ErrorDTO{
		Message: err.Error(),
		Time:    time.Now(),
	}
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "	")

	if err != nil {
		panic(err)
	}

	return string(b)
}
