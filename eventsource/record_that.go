package eventsource

import (
	"fmt"
	"github.com/theshaunwalker/eventsourcing/util"
	"reflect"
)

type RecordThatContainer struct {
	event Event
}

func RecordThat(event Event) RecordThatContainer {
	return RecordThatContainer{event: event}
}

func (recordThatContainer RecordThatContainer) HappenedOn(root AggregateRoot) {

	root.ApplyEvent(recordThatContainer.event)

	fmt.Println("aa ffs")
	fmt.Println(util.GetTypeName(root))
	nextVersion := root.GetAggregateRootVersion() + 1

	newEventEnvelope := EventEnvelope{
		ID:        0,
		AggId:     root.GetAggregateRootId(),
		AggType:   util.GetTypeName(root),
		AggIndex:  nextVersion,
		EventName: reflect.TypeOf(recordThatContainer.event).Name(),
		Event:     recordThatContainer.event,
	}

	fmt.Println("New event envelope")
	fmt.Printf("%v\n", newEventEnvelope)
	StoreEvent(newEventEnvelope)
}
