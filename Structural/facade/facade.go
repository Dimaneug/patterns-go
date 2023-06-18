package main

import "fmt"

type Computer struct{}

func (c *Computer) GetElectricShock() {
	fmt.Println("Current appeared")
}

func (c *Computer) MakeSound() {
	fmt.Println("Beep")
}

func (c *Computer) ShowLoadingScreen() {
	fmt.Println("Loading")
}

func (c *Computer) ShowDesktop() {
	fmt.Println("System is ready")
}

func (c *Computer) CloseEverything() {
	fmt.Println("Closed all apps")
}

func (c *Computer) StopElectricShock() {
	fmt.Println("Current disappeared")
}

type ComputerFacade struct {
	computer Computer
}

func (cf *ComputerFacade) TurnOn() {
	cf.computer.GetElectricShock()
	cf.computer.MakeSound()
	cf.computer.ShowLoadingScreen()
	cf.computer.ShowDesktop()
	fmt.Println("Turned on")
}

func (cf *ComputerFacade) TurnOff() {
	cf.computer.ShowLoadingScreen()
	cf.computer.CloseEverything()
	cf.computer.StopElectricShock()
	fmt.Println("Turned off")
}

func main() {
	computer := ComputerFacade{}
	computer.TurnOn()
	computer.TurnOff()
}
