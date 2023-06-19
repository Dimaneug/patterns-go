package main

import (
	"fmt"
	"strings"
)

type WritingState interface {
	Write(words string)
}

type UpperCase struct{}

func (uc *UpperCase) Write(words string) {
	fmt.Println(strings.ToUpper(words))
}

type LowerCase struct{}

func (lc *LowerCase) Write(words string) {
	fmt.Println(strings.ToLower(words))
}

type Default struct{}

func (d *Default) Write(words string) {
	fmt.Println(words)
}

type TextEditor struct {
	state WritingState
}

func (te *TextEditor) SetState(state WritingState) {
	te.state = state
}

func (te *TextEditor) TypeWords(words string) {
	te.state.Write(words)
}

func main() {
	editor := TextEditor{state: &Default{}}
	editor.TypeWords("Firstr string")

	editor.SetState(&UpperCase{})
	editor.TypeWords("Second string")

	editor.SetState(&LowerCase{})
	editor.TypeWords("Third string")
}
