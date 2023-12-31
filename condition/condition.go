package main

import "fmt"

var score int
func grade() {
	fmt.Println("Grade calculator")
	fmt.Scanf("%d", &score)
	fmt.Println(" score = ", score)
	if score >= 80 {
		fmt.Println("Grade is A")
	}else if score >= 70{
		fmt.Println("Grade is B")
	}else if score >= 60{
		fmt.Println("Grade is C")
	}else if score >= 50{
		fmt.Println("Grade is D")
	}else {
		fmt.Println("Grade is F")
	}
}

func main() {
	grade()
}