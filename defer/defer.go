package main

import (
	"fmt"
)

func add(value1, value2 float64) {
	result := value1 + value2
	fmt.Println("result : ", result)
}

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println("i = ", i)
	}
}

func deferloop() {
	for j := 0; j < 10; j++ {
		defer fmt.Println("j = ", j)
	}
}

func main() {
	loop()
	deferloop()

	// fmt.Println("Calculator !!")
	
	// defer last in first out
	// 14 20 30 End
	// defer fmt.Println("End")
	// defer add(20, 10)
	// defer add(15, 5)
	// defer add(12, 2)

}