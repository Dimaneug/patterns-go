package main

import (
	"fmt"
	"sync"
)

type Wolf struct {
	Name string
	Age  uint8
}

var (
	once     sync.Once
	instance singleton
)

type singleton Wolf

func New(name string, age uint8) singleton {
	once.Do(func() {
		instance = singleton{Name: name, Age: age}
	})

	return instance
}

func main() {
	wolf := New("Oleg", 100)
	fakeWolf := New("Alex", 25)
	fmt.Println(wolf)
	fmt.Println(fakeWolf)

}
