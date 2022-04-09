package main

import "fmt"

func addHundred(num int) {
  num += 100
}

func brainwash(saying *string) { 
	*saying = "Beep Boop."
}

func oppositeDay(boolean *bool) {
	if *boolean == true {
		*boolean = false
	} else {
		*boolean = true
	}
}

func interestRate(percent *float64) {
	*percent = *percent * 1.2
}

func main() {
  x := 1
  
  addHundred(x)
  fmt.Println(x) 

  	treasure := "The friends we make along the way."
	fmt.Println(&treasure)

	star := "Polaris"
	starAddress := &star
	fmt.Println("The address of star is", starAddress)
	*starAddress = "Sirius"
	fmt.Println("The actual value of star is", star)

	greeting := "Hello there!"
	brainwash(&greeting)
	fmt.Println("greeting is now:", greeting)

	trueVar := true
	oppositeDay(&trueVar)
	fmt.Println("true is now:", trueVar)

	currentInterestRate := 5.6
	interestRate(&currentInterestRate)
	fmt.Println("the new interest rate is:", currentInterestRate)
}