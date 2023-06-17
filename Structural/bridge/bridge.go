package main

import "fmt"

type Theme interface {
	GetColor() string
}

type WebPage interface {
	GetContent() string
}

type About struct {
	theme Theme
}

func (a *About) GetContent() string {
	return "About page with " + a.theme.GetColor()
}

type Projects struct {
	theme Theme
}

func (a *Projects) GetContent() string {
	return "Projects page with " + a.theme.GetColor()
}

type Careers struct {
	theme Theme
}

func (a *Careers) GetContent() string {
	return "Careers page with " + a.theme.GetColor()
}

type DarkTheme struct{}

func (dt *DarkTheme) GetColor() string {
	return "dark theme"
}

type LightTheme struct{}

func (dt *LightTheme) GetColor() string {
	return "light theme"
}

type AquaTheme struct{}

func (dt *AquaTheme) GetColor() string {
	return "aqua theme"
}

func main() {
	darkTheme := DarkTheme{}

	about := About{theme: &darkTheme}
	careers := Careers{theme: &darkTheme}

	fmt.Println(about.GetContent())
	fmt.Println(careers.GetContent())
}
