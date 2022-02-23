package main

import "fmt"

func main() {

	x := 123
	y := &x
	var z *int = &x
	var w = myPointer(&x)

	fmt.Println(&x) // memory address
	fmt.Println(*y) // value of memory
	fmt.Println(z)  // memory address
	fmt.Println(*z) // value of memory
	fmt.Println(w)  // value of memory

}

func myPointer(i *int) int {
	return *i
}
