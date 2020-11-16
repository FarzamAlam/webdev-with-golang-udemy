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

func (p person) pSpeak() {
	fmt.Println(p.name, " speaks.")
}

func (sa secretAgent) saSpeak() {
	fmt.Println(sa.p.name, " agent speaks.")
}

func main() {
	p := person{
		name:   "Farzam Alam",
		gender: "Male",
	}
	sa := secretAgent{
		p:     p,
		level: "11",
	}
	p.pSpeak()
	fmt.Println("Gender of person : ", p.gender)

	sa.saSpeak()
	fmt.Println("Level of secret agent : ", sa.level)
}
