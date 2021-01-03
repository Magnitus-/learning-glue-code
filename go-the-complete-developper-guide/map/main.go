package main

import "fmt"

func printMap(m map[string]string) {
	for key, val := range m {
		fmt.Println(key, ":", val)
	}
}

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"black": "#000000",
	}
	var colorsEmpty map[string]string
	colorsEmptyAlternative := make(map[string]string)

	colors["white"] = "#ffffff"
	delete(colors, "green")

	fmt.Println(colors)
	fmt.Println(colorsEmpty)
	fmt.Println(colorsEmptyAlternative)

	printMap(colors)
}
