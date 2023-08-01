package dummyapp

import (
	"fmt"
	"github.com/theshaunwalker/eventsourcing/dummyapp/users"
	"github.com/theshaunwalker/eventsourcing/eventsource"
)

func Sandbox() {
	user := users.UserSignedUp(
		"Shaun",
		"shaun@example.com",
	)

	eventsource.StoreAggregateRoot(user)

	fmt.Println("Done!")
}
