package main

import (
	"fmt"
	"math/rand"
	"math"
	"math/cmplx"
	"time"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java = true, false, "no!"

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

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

	fmt.Println(split(29))

	var i int
	fmt.Println(i, c, python, java)

	var j, k int = 1, 2
	l := 3
	fmt.Println(j, k, l)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var s string
	fmt.Printf("%q\n", s)

	v := 0.867 + 0.5i
	fmt.Printf("v is of type %T\n", v)

	const World = "世界"
	fmt.Println("Hello", World)
}