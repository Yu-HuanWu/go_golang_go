package main

import "fmt"

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
}
