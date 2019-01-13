package main

import "fmt"

type englishBot struct {}
type spanishBot struct {}

type bot interface {
	getGreeting()string
}

func main(){
	eb := englishBot{}
	sp := spanishBot{}

	printGreeting(eb)
	printGreeting(sp)
}

func (eb englishBot) getGreeting() string{
	return "Hi there!"
}
func (sp spanishBot) getGreeting() string{
	return "Hola!"
}
func printGreeting(b bot){
	fmt.Println(b.getGreeting())
}

