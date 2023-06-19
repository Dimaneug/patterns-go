package main

import "fmt"

type Builerable interface {
	Test()
	Lint()
	Assemble()
	Deploy()
	Build(i Builerable)
}

type Builder struct{}

func (b *Builder) Build(i Builerable) {
	i.Test()
	i.Lint()
	i.Assemble()
	i.Deploy()
}

type AndroidBuilder struct {
	Builder
}

func (ab *AndroidBuilder) Test() {
	fmt.Println("Execute android tests")
}

func (ab *AndroidBuilder) Lint() {
	fmt.Println("Execute android linter")
}

func (ab *AndroidBuilder) Assemble() {
	fmt.Println("Building android app")
}

func (ab *AndroidBuilder) Deploy() {
	fmt.Println("Deploying android build on server")
}

type IOSBuilder struct {
	Builder
}

func (ib *IOSBuilder) Test() {
	fmt.Println("Execute ios tests")
}

func (ib *IOSBuilder) Lint() {
	fmt.Println("Execute ios linter")
}

func (ib *IOSBuilder) Assemble() {
	fmt.Println("Building ios app")
}

func (ib *IOSBuilder) Deploy() {
	fmt.Println("Deploying ios build on server")
}

func main() {
	androidBuilder := AndroidBuilder{}
	androidBuilder.Build(&androidBuilder)

	iosBuilder := IOSBuilder{}
	iosBuilder.Build(&iosBuilder)
}
