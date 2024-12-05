package main

import "fmt"

type humantype interface {
	speak()
}

type PersonType struct {
	name string
}

func (p *PersonType) speak() {
	fmt.Println("My name is ", p.name, "and I speak")
}

func saySum(h humantype) {
	h.speak()
}
