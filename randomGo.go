package main

import (
	"fmt"
	"math/rand"
	"math"
	"time"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
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

	a, b := swap("hello", "world")
	fmt.Println(a, b)
}