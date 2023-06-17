package main

import "fmt"

type Sheep struct {
	name     string
	category string
}

func main() {
	originalSheep := Sheep{
		name:     "Djolli",
		category: "Common sheep",
	}
	cloneSheep := originalSheep
	fmt.Println(originalSheep == cloneSheep)
	fmt.Println(&originalSheep == &cloneSheep)

}
