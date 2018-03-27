package main

import "fmt"

func main() {
	///strictly typed key value pair

	//var colors map[string]string

	//colors := make(map[string]string)

	//colors["white"] = "fff56e"

	//delete(colors, "white")

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4324f54",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
