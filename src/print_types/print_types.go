package main

import "fmt"

func main() {

	myName := "My Name"
	myAge := 20
	myBill := (5.0 / 1025)
	myBoolean := false

	fmt.Printf("Value = `%v` ", myName)
	fmt.Printf("Type = `%T` \n", myName)

	fmt.Printf("Value = `%v` ", myAge)
	fmt.Printf("Type = `%T` \n", myAge)

	fmt.Printf("Value = `%v` ", myBill)
	fmt.Printf("Type = `%T` \n", myBill)

	fmt.Printf("Value = `%v` ", myBoolean)
	fmt.Printf("Type = `%T` \n", myBoolean)

}
