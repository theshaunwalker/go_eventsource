package users

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/theshaunwalker/eventsourcing/eventsource"
	"github.com/theshaunwalker/eventsourcing/util"
	"reflect"
)

type User struct {
	eventsource.AggregateRootBehavior
	id    uuid.UUID
	name  string
	email string
}

func (user User) GetAggregateRootId() uuid.UUID {
	return user.id
}

func (user *User) ApplyEvent(event eventsource.Event) {
	theType := reflect.TypeOf(event).Name()
	methodName := "Apply" + theType

	//fmt.Println("user ApplyEvent")
	//fmt.Printf("%v\n", theType)

	util.CallDynamicFunctionOnStructInstance(user, methodName, event)
}

func UserSignedUp(name string, email string) User {
	userSignedUp := NewUserSignedUp{
		EventId: uuid.New(),
		UserId:  uuid.New(),
		Name:    name,
		Email:   email,
	}
	user := User{}

	eventsource.RecordThat(userSignedUp).HappenedOn(&user)

	return user
}

func (user *User) ApplyNewUserSignedUp(event NewUserSignedUp) {
	user.id = event.UserId
	user.name = event.Name
	user.email = event.Email

	fmt.Println("user ID")
	fmt.Printf("%v\n", user.id)
}
