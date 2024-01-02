package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(promt string) float64 {
	// %v ใช้ระบุ data type ที่ยังไม่แน่นอน
	fmt.Printf("%v", promt)
	// รับค่า input มาเป็น string
	input, _ := reader.ReadString('\n')
	// แปลง string เป็น float
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message, _ := fmt.Scanf("%v must number only", promt)
		panic(message)
	}
	return value
}

func add(value1, value2 float64) float64 {
	return value1 + value2
}

func minus(value1, value2 float64) float64 {
	return value1 - value2
}

func multiply(value1, value2 float64) float64 {
	return value1 * value2
}

func divide(value1, value2 float64) float64 {
	return value1 / value2
}

func getOperator() string{
	fmt.Print("operator is ( + - * / ):")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func main() {
	numOfInputs := getInput("คุณต้องการคำนวณกี่ตัวเลข : ")
	var result float64

	if numOfInputs < 1 {
		panic("Number of calculations should be at least 1")
	}

	result = getInput("Number 1: ")

	for i := 1; i < int(numOfInputs); i++ {
		operator := getOperator()
		value := getInput(fmt.Sprintf("Number %d: ", i+1))
		
		switch operator {
		case "+":
			result = add(result, value)
		case "-":
			result = minus(result, value)
		case "*":
			result = multiply(result, value)
		case "/":
			result = divide(result, value)
		default:
			panic("wrong operator")
		}
	}

	fmt.Printf("result = %v", result)
}