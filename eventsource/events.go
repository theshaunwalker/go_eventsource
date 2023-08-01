package eventsource

import (
	"encoding/json"
	"github.com/google/uuid"
	"reflect"
)

type Event interface {
	GetEventId() uuid.UUID
	GetJsonBody() string
}

type EventEnvelope struct {
	ID        uint
	AggId     uuid.UUID
	AggType   string
	AggIndex  uint
	EventName string
	Event     Event
}

type Tabler interface {
	TableName() string
}

func (EventEnvelopeStoreRecord) TableName() string {
	return "event_store"
}

type EventEnvelopeStoreRecord struct {
	ID        uint
	AggId     uuid.UUID
	AggType   string
	AggIndex  uint
	EventName string
	Event     string
}

// EventEnvelope <-> EventEnvelopeStoreRecord
func (envelope EventEnvelope) ToStoreRecord() EventEnvelopeStoreRecord {
	jsonEventBody, err := json.Marshal(envelope.Event)

	if err != nil {
		panic("failed to json encode the event body of " + reflect.TypeOf(envelope.Event).Name())
	}
	return EventEnvelopeStoreRecord{
		ID:        envelope.ID,
		AggId:     envelope.AggId,
		AggType:   envelope.AggType,
		AggIndex:  envelope.AggIndex,
		EventName: envelope.EventName,
		Event:     string(jsonEventBody),
	}
}

//func (storeRecord EventEnvelopeStoreRecord) ToEventEnvelope() EventEnvelope {
//
//
//	return EventEnvelope{
//		ID:        storeRecord.ID,
//		AggId:     storeRecord.AggId,
//		AggType:   storeRecord.AggType,
//		AggIndex:  storeRecord.AggIndex,
//		EventName: storeRecord.EventName,
//		Event:     event,
//	}
//}
