package main

import "fmt"

type Person struct {
	Name string
	Age  int32
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type ByAge []Person

func (p ByAge) Len() int { return len(p) }

func (p ByAge) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (p ByAge) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}
