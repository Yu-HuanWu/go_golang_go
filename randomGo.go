package main

import (
	"fmt"
	"math/rand"
	"math"
	"time"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	for i := 0; i < 5; i++ {
		dur := time.Duration(rand.Intn(1000)) * time.Millisecond
		fmt.Printf("Sleeping for %v\n", dur)
		// Sleep for a random duration between 0-1000ms
		time.Sleep(dur)
	}
	fmt.Println("Done!")

	fmt.Println("The time is", time.Now())
	fmt.Println(math.Pi)
	fmt.Println(add(1337, 42))
}