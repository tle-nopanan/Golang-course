package main

import "os"

func check(e error) {
	
	if e != nil {
		panic(e)
	}

}

func main() {
	data1 := []byte("hello\n Tle")
	err := os.WriteFile("data.txt", data1, 0644)
	check(err)
	f, err := os.Create("emplyeeName")
	check(err)

	defer f.Close()

	data2 := []byte("Nopanan\n Manee")
	os.WriteFile("emplyeeName.txt", data2, 0644)
}