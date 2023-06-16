package main

import (
	"fmt"
)

type Warrior interface {
	UseSword()
}

type Archer interface {
	UseBow()
}

type OrcWarrior struct{}

func (ow *OrcWarrior) UseSword() {
	fmt.Println("Orc attacked with sword")
}

type OrcArcher struct{}

func (oa *OrcArcher) UseBow() {
	fmt.Println("Orc attacked with Bow")
}

type ElfWarrior struct{}

func (ew *ElfWarrior) UseSword() {
	fmt.Println("Elf attacked with sword")
}

type ElfArcher struct{}

func (ea *ElfArcher) UseBow() {
	fmt.Println("Elf attacked with Bow")
}

type AbstractFactory interface {
	CreateWarrior() Warrior
	CreateArcher() Archer
}

type OrcFactory struct{}

func (of *OrcFactory) CreateWarrior() Warrior {
	return &OrcWarrior{}
}

func (of *OrcFactory) CreateArcher() Archer {
	return &OrcArcher{}
}

type ElfFactory struct{}

func (of *ElfFactory) CreateWarrior() Warrior {
	return &ElfWarrior{}
}

func (of *ElfFactory) CreateArcher() Archer {
	return &ElfArcher{}
}

func armyAttack(factory AbstractFactory) {
	warrior := factory.CreateWarrior()
	warrior.UseSword()

	archer := factory.CreateArcher()
	archer.UseBow()
}

func main() {
	barrack := OrcFactory{}
	treeOfWisdom := ElfFactory{}

	armyAttack(&barrack)
	armyAttack(&treeOfWisdom)
}
