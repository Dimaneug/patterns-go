package main

import (
	"fmt"
	"strings"
)

type Burger struct {
	kilocalories uint16
	weight       uint16

	cheese    uint8
	tomato    bool
	pickle    bool
	onion     bool
	bacon     bool
	beefPatty uint8
}

func (b *Burger) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf(
		"Burger:\n"+
			"Cal: %d\n"+
			"Weight: %d\n",
		b.kilocalories,
		b.weight,
	))

	if b.cheese > 0 {
		builder.WriteString(fmt.Sprintf("Cheese: %d\n", b.cheese))
	}

	if b.tomato {
		builder.WriteString("Tomato\n")
	}

	if b.pickle {
		builder.WriteString("Pickle\n")
	}

	if b.onion {
		builder.WriteString("Onion\n")
	}

	if b.bacon {
		builder.WriteString("Bacon\n")
	}

	if b.tomato {
		builder.WriteString(fmt.Sprintf("Beef Patty: %d", b.beefPatty))
	}

	return builder.String()

}

// func NewBurger(bb BurgerBuilder) *Burger {}

type BurgerBuilder struct {
	Burger
}

func (bb *BurgerBuilder) AddCheese() *BurgerBuilder {
	bb.cheese++
	bb.kilocalories += 70
	bb.weight += 30
	return bb
}

func (bb *BurgerBuilder) AddTomato() *BurgerBuilder {
	bb.tomato = true
	bb.kilocalories += 20
	bb.weight += 50
	return bb
}

func (bb *BurgerBuilder) AddPickle() *BurgerBuilder {
	bb.pickle = true
	bb.kilocalories += 20
	bb.weight += 50
	return bb
}

func (bb *BurgerBuilder) AddOnion() *BurgerBuilder {
	bb.onion = true
	bb.kilocalories += 25
	bb.weight += 15
	return bb
}

func (bb *BurgerBuilder) AddBacon() *BurgerBuilder {
	bb.bacon = true
	bb.kilocalories += 150
	bb.weight += 80
	return bb
}

func (bb *BurgerBuilder) AddBeefPatty() *BurgerBuilder {
	bb.beefPatty++
	bb.kilocalories += 250
	bb.weight += 130
	return bb
}

func (bb *BurgerBuilder) GetBurger() *Burger {
	return &Burger{
		kilocalories: 50 + bb.kilocalories,
		weight:       50 + bb.weight,
		cheese:       bb.cheese,
		tomato:       bb.tomato,
		pickle:       bb.pickle,
		onion:        bb.onion,
		bacon:        bb.bacon,
		beefPatty:    bb.beefPatty,
	}
}

// func NewBigMac(bb BurgerBuilder) *Burger {
// 	bigmac := bb.
// }

type BurgerDirector struct {
	Builder *BurgerBuilder
}

func (bd *BurgerDirector) Build() {
	bd.Builder.
		AddBeefPatty().
		AddBacon().
		AddCheese().
		AddOnion().
		AddTomato().
		AddPickle()
}

func main() {
	builder := &BurgerBuilder{}
	director := &BurgerDirector{
		Builder: builder,
	}
	director.Build()
	burger := builder.GetBurger()

	fmt.Println(burger)

}
