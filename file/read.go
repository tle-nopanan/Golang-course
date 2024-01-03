package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("exam.csv")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	count := 1
	for scanner.Scan(){
		line := scanner.Text()
		item := strings.Split(line, ",")
		fmt.Println(item[1])
		// fmt.Printf("Line %d : %s\n", count, line)
		
		count++
	}

}