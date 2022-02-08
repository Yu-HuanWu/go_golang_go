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
}
