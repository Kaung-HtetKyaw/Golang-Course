package main

import "fmt"

func AddT[T int | float64](a, b T) T {
	return a + b
}

type Walkable interface {
	Walk()
}

type Duck struct {
	name string
}

func (d Duck) Walk() {
	fmt.Println("My name is ", d.name, "and I can walk")
}

type WalkableCreature interface {
	Walkable
}

func getWalkableCreature(creature WalkableCreature) {
	fmt.Println("The creature walks and it says: ")
	creature.Walk()
}
