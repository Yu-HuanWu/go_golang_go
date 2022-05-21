package main

import (
	"fmt"
	"math/rand"
	"math"
	"math/cmplx"
	"time"
	"runtime"
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

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

type Vertex struct {
	X int
	Y int
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

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	factorial := 1
	for i := 1; i < 10; i++ {
		factorial *= i
	}
	fmt.Println(factorial)

	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)

	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	switch {
	case t.Hour() == 12:
		fmt.Println("Lunch Time!")
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	i1, j1 := 42, 2701

	p := &i1         // point to i1
	fmt.Println(*p) // read i1 through the pointer
	*p = 21         // set i1 through the pointer
	fmt.Println(i1)  // see the new value of i1

	p = &j1         // point to j1
	*p = *p / 37   // divide j1 through the pointer
	fmt.Println(j1) // see the new value of j1

	fmt.Println(Vertex{1, 2})
	ver := Vertex{1, 2}
	ver.X = 4
	fmt.Println(ver.X)

	p2 := &ver
	p2.X = 1e9
	fmt.Println(ver)

	//.. last
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")

	defer fmt.Println("world")
	fmt.Println("hello")
}