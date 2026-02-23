package main

import "fmt"

type colorMap map[string]string

func main() {
	colors := colorMap{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
		"white": "#FFFFFF",
	}

	printMap(colors)
}

func printMap(c colorMap) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
