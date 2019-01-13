package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastNme string
	contact contactInfo
}

func main() {

	alex := person{"Joseph", "Attakorah", contactInfo{"juxluvjoe@gmail.com", 15}}
	fmt.Println(alex)
	alex.firstName = "John"
	fmt.Println(alex)
	alex.updateName("Hilda")
	fmt.Println(alex)
	fmt.Println(alex.contact.email)
}

func (p *person) updateName(name string){
	(*p).firstName = name
}
