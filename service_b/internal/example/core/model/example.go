package model

import (
	"encoding/json"
	"time"

	"github.com/oklog/ulid/v2"
)

type Example struct {
	ID        string     `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"iupdated_atd"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

type NewExampleInput struct {
	Name string
}

func NewExample(in NewExampleInput) *Example {
	return &Example{
		ID:        ulid.Make().String(),
		Name:      in.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}
}

func (m *Example) MarshalJSON() ([]byte, error) {
	copy := *m

	return json.Marshal(copy)
}
