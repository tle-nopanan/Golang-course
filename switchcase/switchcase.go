package main

import "fmt"

func main() {
	var input string
	fmt.Println("Input Color")
	fmt.Scanf("%s", &input)
	switch input {
	case "blue":
		fmt.Println(input," = #0000FF")
	case "green":
		fmt.Println(input," = #008000")
	case "pink":
		fmt.Println(input," = #FFC0CB")
	case "yellow":
		fmt.Println(input," = #FFFF00")
	default:
		fmt.Println(input," = invalid color")
	}

	// switch input {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// default:
	// 	fmt.Println("invalid value")
	// }
}