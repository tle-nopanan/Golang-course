package main

import "fmt"

func main() {
	var courseName []string
	courseName = []string{"Java", "Python"}
	fmt.Println(courseName)
	courseName = append(courseName, "C", "C#", "HTML", "CSS", "JacaScript")
	fmt.Println(courseName)

	courseWeb := courseName[4:7]
	fmt.Println(courseWeb)
	
	courseWeb = courseName[0:4]
	fmt.Println(courseWeb)
}