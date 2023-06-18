package main

import "fmt"

type Coffee struct {
	kind string
}

type Barista struct {
	availableCoffee map[string]*Coffee
}

func (b *Barista) MakeCoffee(kind string) *Coffee {
	if _, found := b.availableCoffee[kind]; !found {
		fmt.Println("Created new coffee", kind)
		coffee := Coffee{kind: kind}
		b.availableCoffee[kind] = &coffee
		return &coffee
	}
	fmt.Println("Return previously cooked coffee", kind)
	return b.availableCoffee[kind]

}

func MakeBarista() *Barista {
	return &Barista{availableCoffee: make(map[string]*Coffee)}
}

type CoffeeShop struct {
	orders  []*Coffee
	barista *Barista
}

func (cs *CoffeeShop) TakeOrder(kind string) {
	coffee := cs.barista.MakeCoffee(kind)
	cs.orders = append(cs.orders, coffee)
}

func (cs *CoffeeShop) ShowOrders() {
	for i, val := range cs.orders {
		fmt.Println(i, "-", *val)
	}
}

func main() {
	boris := MakeBarista()
	coffeeShop := CoffeeShop{barista: boris}
	coffeeShop.TakeOrder(("americano"))
	coffeeShop.TakeOrder(("cappucino"))
	coffeeShop.TakeOrder(("americano"))

	coffeeShop.ShowOrders()

}
