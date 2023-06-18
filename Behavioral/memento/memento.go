package main

import "fmt"

type EditorMemento struct {
	content string
}

func (em *EditorMemento) GetContent() string {
	return em.content
}

type Editor struct {
	content string
	backup  EditorMemento
}

func (e *Editor) TypeText(text string) {
	e.content += text
}

func (e *Editor) Save() {
	e.backup.content = e.content
}

func (e *Editor) Restore() {
	e.content = e.backup.content
}

func main() {
	editor := Editor{backup: EditorMemento{}}
	editor.TypeText("Hello")
	editor.Save()
	editor.TypeText(" world")
	fmt.Println(editor.content)
	editor.Restore()
	fmt.Println(editor.content)

}
