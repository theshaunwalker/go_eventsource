package eventsource

import (
	"github.com/google/uuid"
)

type AggregateRoot interface {
	GetAggregateRootId() uuid.UUID
	GetAggregateRootVersion() uint
	ApplyEvent(event Event)
}

type AggregateRootBehavior struct {
	EventRecording
	AggregateRootVersion uint
}

type EventRecording struct {
	events []EventEnvelope
}

type RecordsEvents interface {
	ReleaseEvents() []EventEnvelope
}

func (recordsEvents EventRecording) ReleaseEvents() []EventEnvelope {
	releasedEvents := recordsEvents.events

	recordsEvents.events = make([]EventEnvelope, 0)

	return releasedEvents
}

func (root AggregateRootBehavior) GetAggregateRootVersion() uint {
	return root.AggregateRootVersion
}
