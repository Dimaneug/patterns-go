package main

import "fmt"

type Animal interface {
	Accept(operation AnimalOperation)
}

type AnimalOperation interface {
	VisitMonkey(monkey Monkey)
	VisitLion(lion Lion)
	VisitDolphin(dolphin Dolphin)
}

type Monkey struct{}

func (m *Monkey) Shout() {
	fmt.Println("Monkey is shouting")
}

func (m *Monkey) Accept(operation AnimalOperation) {
	operation.VisitMonkey(*m)
}

type Lion struct{}

func (l *Lion) Roar() {
	fmt.Println("Lion is roaring")
}

func (l *Lion) Accept(operation AnimalOperation) {
	operation.VisitLion(*l)
}

type Dolphin struct{}

func (d *Dolphin) Speak() {
	fmt.Println("Dolphin is speaking")
}

func (d *Dolphin) Accept(operation AnimalOperation) {
	operation.VisitDolphin(*d)
}

type Visitor struct{}

func (v *Visitor) VisitMonkey(monkey Monkey) {
	monkey.Shout()
}

func (v *Visitor) VisitLion(lion Lion) {
	lion.Roar()
}

func (v *Visitor) VisitDolphin(dolphin Dolphin) {
	dolphin.Speak()
}

func main() {
	monkey := Monkey{}
	lion := Lion{}
	dolphin := Dolphin{}

	visitor := Visitor{}

	monkey.Accept(&visitor)
	lion.Accept(&visitor)
	dolphin.Accept(&visitor)
}
