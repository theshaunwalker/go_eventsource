package users

import (
	"encoding/json"
	"github.com/google/uuid"
)

type NewUserSignedUp struct {
	EventId uuid.UUID
	UserId  uuid.UUID
	Name    string
	Email   string
}

func (event NewUserSignedUp) GetEventId() uuid.UUID {
	return event.EventId
}

func (event NewUserSignedUp) GetJsonBody() string {
	eventJson, _ := json.Marshal(event)
	return string(eventJson)
}
