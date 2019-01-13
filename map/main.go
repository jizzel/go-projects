package main

import "fmt"

func main() {
	colors := map[string] string {
		"red": "#ff0000",
		"green": "#A7C29",
		"white": "#fff",
	}

	//var colors map[string]string

	//colors := make(map[string]string)

	//colors["white"] = "#fff"
	//colors["black"] = "#000"

	//delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
