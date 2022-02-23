package main

import "fmt"

func main() {
	simpleArray()
	slice()
	simpleMap()
}

func simpleArray() {
	myArray := [2]int{1, 2}
	fmt.Println(myArray)
}

func slice() {
	slice := make([]int, 2)
	slice[0] = 1
	slice[1] = 2
	slice = append(slice, 3, 4, 5)
	fmt.Println(slice)
}

func simpleMap() {
	myMap := make(map[string]int)
	myMap["Hello"] = 10
	myMap["World"] = 20
	fmt.Println(myMap["Hello"])
}
