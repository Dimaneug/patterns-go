package main

import "fmt"

type Lamp struct{}

func (l *Lamp) TurnOn() {
	fmt.Println("Lamp switched on")
}

func (l *Lamp) TurnOff() {
	fmt.Println("Lamp switched off")
}

type Command interface {
	Execute()
	Undo()
	Redo()
}

type TurnOn struct {
	lamp Lamp
}

func (to *TurnOn) Execute() {
	to.lamp.TurnOn()
}

func (to *TurnOn) Undo() {
	to.lamp.TurnOff()
}

func (to *TurnOn) Redo() {
	to.lamp.TurnOn()
}

type TurnOff struct {
	lamp Lamp
}

func (to *TurnOff) Execute() {
	to.lamp.TurnOff()
}

func (to *TurnOff) Undo() {
	to.lamp.TurnOn()
}

func (to *TurnOff) Redo() {
	to.lamp.TurnOff()
}

type RemoteControl struct{}

func (rc *RemoteControl) Submit(command Command) {
	command.Execute()
}

func main() {
	lamp := Lamp{}

	turnOn := TurnOn{lamp}
	turnOff := TurnOff{lamp}

	remote := RemoteControl{}
	remote.Submit(&turnOn)
	remote.Submit(&turnOff)
}
