package main

import "fmt"

type something interface {
	showName()
}

type SomeShit struct {
	name string
}

func (s *SomeShit) showName() {
	fmt.Println(s.name)
}

func showSomeShitName(s something) {
	s.showName()
}
