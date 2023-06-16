package main

import "fmt"

type Door interface {
	GetWidth() float32
	GetHeight() float32
}

type WoodenDoor struct {
	width  float32
	height float32
}

func (wd *WoodenDoor) NewWoodenDoor(width, height float32) Door {
	return &WoodenDoor{
		width:  width,
		height: height,
	}
}

func (wd *WoodenDoor) GetWidth() float32 { return wd.width }

func (wd *WoodenDoor) GetHeight() float32 { return wd.height }

type DoorFactory struct{}

func (df *DoorFactory) MakeDoor(width, height float32) Door {
	return &WoodenDoor{
		width:  width,
		height: height,
	}
}

func main() {
	doorFactory := DoorFactory{}
	door := doorFactory.MakeDoor(100, 200)
	fmt.Println(door.GetWidth())
	fmt.Println(door.GetHeight())

}
