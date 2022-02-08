package main

import "fmt"

func main() {
	var world string = "Hello World"
	var happy string = "happy"
	const year = 2022
	var excla string = "!"

	message := fmt.Sprintf("%s, %s %d%s", world, happy, year, excla)
	fmt.Printf(message)
}
