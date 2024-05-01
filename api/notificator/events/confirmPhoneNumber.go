package events

import "github.com/google/uuid"

type ConfirmPhoneNumber struct {
	EventId uuid.UUID
	UserId  uint
	Code    string
}
