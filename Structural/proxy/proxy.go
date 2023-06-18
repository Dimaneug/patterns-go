package main

import "fmt"

type Door interface {
	Open()
	Close()
}

type LabDoor struct{}

func (ld *LabDoor) Open() {
	fmt.Println("Lab door opened")
}

func (ld *LabDoor) Close() {
	fmt.Println("Lab door closed")
}

type Security struct {
	door Door
}

func (s *Security) authenticate(password string) bool {
	return password == "secret"
}

func (s *Security) Open(password string) {
	if s.authenticate(password) {
		s.door.Open()
	} else {
		fmt.Println("Door didn't open")
	}

}

func (s *Security) Close() {
	s.door.Close()
}

func main() {
	door := Security{door: &LabDoor{}}
	door.Open("invalid")

	door.Open("secret")
	door.Close()

}
