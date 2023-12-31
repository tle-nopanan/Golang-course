package main

import "fmt"

func zerovalue(ivalue int) {
	ivalue = 0
}

func zeropointer(ipointer *int) {
	*ipointer = 0
}

func main() {
	i := 1
	fmt.Println("i = ", i)

	zerovalue(i)
	fmt.Println("i from zerovalue func = ", i)

	zeropointer(&i)
	fmt.Println("i value from zeropointer func = ", i)
	fmt.Println("i address from zeropointer func = ", &i)

}