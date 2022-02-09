package main

import (
	"os"
	"sort"
	"time"
)

const timeOfWatiting = 5
const repeatCount = 10

func main() {
	//slices (in arrays, you must declare the size)
	names := []string{"John", "Anna", "Hellen"}
	// sort names
	sort.Strings(names)
	// print one by one...
	for i, name := range names {
		println(i, name)
	}

	// default for...
	for i := 0; i < 5; i++ {
		print(i)
	}

	// infinite loop with poor brake example
	println("Press Ctrl+C to exit")
	count := repeatCount
	for {
		count--
		println("waiting...")
		if count <= 0 {
			println("The End.")
			os.Exit(0)
		}
		time.Sleep(timeOfWatiting * time.Second)
	}
}
