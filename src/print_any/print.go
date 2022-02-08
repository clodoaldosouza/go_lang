package main

import "fmt"

func main() {
	test()
}

func test() {
	var value int = 20
	message := fmt.Sprintf("%d", value)
	println(message)
}
