package main

import (
	"fmt"
	"time"
)

type User interface {
	GetName() string
	Send(message string)
	Receive(message string)
}

type John struct {
	name     string
	mediator ChatRoomMediator
}

func (j *John) GetName() string {
	return j.name
}

func (j *John) Send(message string) {
	j.mediator.Send(j, message)
}

func (j *John) Receive(message string) {
	fmt.Printf("%s received \"%s\"\n", j.name, message)
}

type Nik struct {
	name     string
	mediator ChatRoomMediator
}

func (n *Nik) GetName() string {
	return n.name
}

func (n *Nik) Send(message string) {
	n.mediator.Send(n, message)
}

func (n *Nik) Receive(message string) {
	fmt.Printf("%s received \"%s\"\n", n.name, message)
}

type ChatRoomMediator interface {
	Send(user User, message string)
}

type ChatRoom struct {
	user1 User
	user2 User
}

func (cr *ChatRoom) Send(user User, message string) {
	time := time.Now()
	fmt.Printf("%s [%s]: %s\n", time.Format("2006/01/02 15:04:05:"), user.GetName(), message)
	if user == cr.user1 {
		cr.user2.Receive(message)
	} else {
		cr.user1.Receive(message)
	}
}

func (cr *ChatRoom) SetFirstUser(user User) {
	cr.user1 = user
}

func (cr *ChatRoom) SetSecondUser(user User) {
	cr.user2 = user
}

func main() {
	mediator := ChatRoom{}

	john := &John{name: "John", mediator: &mediator}
	nik := &Nik{name: "Nik", mediator: &mediator}

	mediator.SetFirstUser(john)
	mediator.SetSecondUser(nik)

	john.Send("Goy")
	nik.Send("Skot")

}
