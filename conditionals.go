package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	heistReady := false
	if (heistReady) {
		fmt.Println("Let's Go!")
	} else {
		fmt.Println("uh-oh")
	}

	lockCombo := "2-35-19"
	robAttempt := "1-1-1"
	if (lockCombo == robAttempt) {
		fmt.Println("The vault is now opened.")
	}

	vaultAmt := 2356468
	if (vaultAmt >= 200000) {
		fmt.Println("We're going to need more bags.")
	}

	rightTime := true
	rightPlace := true
	if rightTime && rightPlace {
		fmt.Println("We're outta here!")
	} else {
		fmt.Println("Be patient...")
	}
	enoughRobbers := false
	enoughBags := true
	if enoughRobbers || enoughBags {
		fmt.Println("Grab everything!")
	} else {
		fmt.Println("Grab whatever you can!")
	}

	readyToGo := true
	if !readyToGo {
		fmt.Println("Start the car!")
	} else {
		fmt.Println("What are we waiting for??")
	}

	amountStolen := 64650
	if amountStolen > 1000000 {
		fmt.Println("We've hit the jackpot!")
	} else if amountStolen >= 5000 {
		fmt.Println("Think of all the candy we can buy!")
	} else {
		fmt.Println("Why did we even do this?")
	}

	name := "H. J. Simp"
	switch name {
	case "Butch":
		fmt.Println("Head to Robbers Roost.")
	case "Bonnie":
		fmt.Println("Stay put in Joplin.")
	default:
		fmt.Println("Just hide!")
	}

	if success := true; success {
		fmt.Println("We're rich!")
	} else {
		fmt.Println("Where did we go wrong?")
	}
	amountStolen2 := 50000
	switch numOfThieves := 3; {
	case numOfThieves==1:
		fmt.Println("I'll take all $", amountStolen2)
	case numOfThieves < 6:
	  fmt.Println("Everyone gets $", amountStolen2/numOfThieves)
	default:
		fmt.Println("There's not enough to go around...")
	}

	rand.Seed(time.Now().UnixNano())
	amountLeft := rand.Intn(10000)
	fmt.Println("amountLeft is: ", amountLeft)
	if amountLeft > 5000 {
		fmt.Println("What should I spend this on?")
	} else {
		fmt.Println("Where did all my money go?")
	}
}
