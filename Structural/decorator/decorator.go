package main

import "fmt"

type Coffee interface {
	GetCost() uint
	GetDescription() string
}

type SimpleCoffee struct{}

func (sc *SimpleCoffee) GetCost() uint {
	return 10
}

func (sc *SimpleCoffee) GetDescription() string {
	return "Coffee"
}

type MilkCoffee struct {
	coffee Coffee
}

func (mc *MilkCoffee) GetCost() uint {
	return mc.coffee.GetCost() + 2
}

func (mc *MilkCoffee) GetDescription() string {
	return mc.coffee.GetDescription() + ", milk"
}

type CreamCoffee struct {
	coffee Coffee
}

func (cc *CreamCoffee) GetCost() uint {
	return cc.coffee.GetCost() + 5
}

func (cc *CreamCoffee) GetDescription() string {
	return cc.coffee.GetDescription() + ", cream"
}

type VanillaCoffee struct {
	coffee Coffee
}

func (vc *VanillaCoffee) GetCost() uint {
	return vc.coffee.GetCost() + 3
}

func (vc *VanillaCoffee) GetDescription() string {
	return vc.coffee.GetDescription() + ", vanilla"
}

func main() {
	someCoffee := SimpleCoffee{}
	fmt.Println(someCoffee.GetCost())
	fmt.Println(someCoffee.GetDescription())

	someMCoffee := MilkCoffee{coffee: &someCoffee}
	fmt.Println(someMCoffee.GetCost())
	fmt.Println(someMCoffee.GetDescription())

	someMCCoffee := CreamCoffee{coffee: &someMCoffee}
	fmt.Println(someMCCoffee.GetCost())
	fmt.Println(someMCCoffee.GetDescription())

	someMCVCoffee := VanillaCoffee{coffee: &someMCCoffee}
	fmt.Println(someMCVCoffee.GetCost())
	fmt.Println(someMCVCoffee.GetDescription())
}
