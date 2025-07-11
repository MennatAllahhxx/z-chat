// This package contains the domain models for the chat application.
package models

import (
	"time"
)

type Message struct {
	ID        string    `json:"id" validate:"uuid4"`
	Sender    string    `json:"sender" validate:"required,uuid4"`
	Receiver  string    `json:"receiver" validate:"required,uuid4,nefield=Sender"`
	Content   string    `json:"content" validate:"required,min=1,max=1000"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *Message) Validate() error {
	return validate.Struct(m)
}
