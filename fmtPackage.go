package main

import "fmt"

func main() {
  // There's a mix of Println and Print!
	fmt.Println("Can", "you", "tell", "the", "difference?")
	fmt.Print("Between", "these", "two", "methods?") //no space
	fmt.Print("Anything odd about", "the spacing? \n")
	fmt.Println("Don't worry if you can't spot it, we'll go through this together!")

	animal1 := "cat"
	animal2 := "dog"
	fmt.Printf("Are you a %v or a %v person?", animal1, animal2)

	fmt.Println("\n***")
	floatExample := 1.75
	fmt.Printf("Working with a %T", floatExample) 
	fmt.Println("\n***") // Added for spacing
	yearsOfExp := 3
	reqYearsExp := 15
	fmt.Printf("I have %d years of Go experience and this job is asking for %d years.", yearsOfExp, reqYearsExp) 
	fmt.Println("\n***") // Added for spacing
	stockPrice := 3.50
	fmt.Printf("Each share of Gopher feed is $%.2f", stockPrice)

	step1 := "Breathe in..."
	step2 := "Breathe out..."
	meditation := fmt.Sprintln(step1, step2)
	fmt.Println(meditation)

	template := "I wish I had a %v."
	pet := "puppy"
	var wish string
}
