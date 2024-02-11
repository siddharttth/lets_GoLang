package main

import (
	"fmt"
)

func main() {
	// var colors map[string]string
	//---------------MAP1------------------
	// colors := make(map[string]string)
	//-------------------------------------
	// colors := map[string]string{
	// 	"red":   "#apple",
	// 	"green": "#guava",
	// }
	//-------------------------------------
	// colors["White"] = "#paper"
	//fmt.Println(colors)
	//----------------MAP2-----------------
	// dhoni := make(map[int]string)
	// dhoni[7] = "thala-for-a-reason"
	//fmt.Println(dhoni)
	//----------------MAP_DEL-------------
	// delete(colors, "White")

	// fmt.Println(colors)
	// fmt.Println(dhoni)
	//-------------------------------------
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#ff2123",
		"white": "#fff112",
	}
	printMap(colors)
}

func printMap(c map[string]string) {
	for key, val := range c {
		fmt.Println("Hex code for ", key, " is ", val)
	}
}
