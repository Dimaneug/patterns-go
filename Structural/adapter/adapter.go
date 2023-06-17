package main

type Lion interface {
	Roar()
}

type AfricanLion struct{}

func (al *AfricanLion) Roar() {}

type AsianLion struct{}

func (al *AsianLion) Roar() {}

type Hunter struct{}

func (h *Hunter) Hunt(lion Lion) {}

type WildDog struct{}

func (wd *WildDog) Bark() {}

type WildDogAdapter struct {
	dog WildDog
}

func (wda *WildDogAdapter) Roar() {
	wda.dog.Bark()
}

func main() {
	wildDog := WildDog{}
	wildDogAdapter := WildDogAdapter{dog: wildDog}

	hunter := Hunter{}
	hunter.Hunt(&wildDogAdapter)
}
