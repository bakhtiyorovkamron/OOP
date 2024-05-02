package main

import "fmt"

type Parent struct{}

func (p *Parent) Print() {
	fmt.Println("Parent")
}

type Child struct {
	Parent
}

func (p *Child) Print() {
	fmt.Println("Child")
}

func main() {
	var x Child
	x.Print()
}
