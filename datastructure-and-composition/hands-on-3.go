package main

import "fmt"

type person struct {
	name   string
	gender string
}

type secretAgent struct {
	p     person
	level string
}

type human interface {
	speak()
}

func (p person) speak() {
	fmt.Println(p.name, "speaks.")
}

func (sa secretAgent) speak() {
	fmt.Println(sa.p.name, " secret agent speaks.")
}

func talk(h human) {
	h.speak()
}

func main() {
	p := person{
		name:   "farzam alam",
		gender: "male",
	}
	sa := secretAgent{
		p:     p,
		level: "11",
	}
	talk(p)
	talk(sa)
}
