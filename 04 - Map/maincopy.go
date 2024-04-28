package main

import "fmt"

func maincopy() {
	//creating an empty map
	//var colors map[string]string
	colors := make(map[string]string)

	/* colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	} */

	colors["white"] = "#ffffff" //adding values

	delete(colors, "white")

	fmt.Println(colors)
}
