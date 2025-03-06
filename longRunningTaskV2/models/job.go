package models

import (
	"time"

	"github.com/google/uuid"
)

// Job represents UUID of a job
type Job struct {
	ID        uuid.UUID `json:"uudi"`
	Type      string    `json:"type"`
	ExtraData any       `json:"extra_data"`
}

// Worker-A data
type Log struct {
	ClientTime time.Time `json:"client_time"`
}

// Callback data
type CallBack struct {
	CallBackURL string `json:"callback_url"`
}

// Mail data
type Mail struct {
	EmailAddress string `json:"email_address"`
}
