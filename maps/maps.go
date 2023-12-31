package main

import "fmt"

var product = make(map[string]float64)

func main() {
	fmt.Println("product = ",product)

	//add
	product["Macbook"] = 40000
	product["iPhone"] = 30000
	product["iPad"] = 25000
	fmt.Println("product = ",product)

	//delete
	delete(product, "iPad")
	fmt.Println("product = ",product)

	//update
	product["iPhone"] = 45000
	fmt.Println("product = ",product)

	//การเข้าถึงข้อมูล
	value1 := product["iPhone"]
	fmt.Println("value1 = ", value1)

	courseName := map[string]string{"101":"Java", "102":"Python"}
	fmt.Println("couseName = ", courseName)
}