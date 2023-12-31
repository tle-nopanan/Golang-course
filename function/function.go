package main

import "fmt"

func hello() {
	fmt.Println("Hello World")
}

func plus(value1 int, value2 int) int {
	return value1 + value2
}

func plus3value(value1,value2,value3 int) int  {
	return value1 + value2 + value3
}

func main() {
	hello()
	result := plus(10,5)
	fmt.Println("result = ", result)
	result2 := plus3value(4,5,11)
	fmt.Println("result2 = ", result2)
}