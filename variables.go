package main
import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(2235*1231)
	const earthsGravity= 9.80665
	var jellybeanCounter int8
	jellybeanCounter +=1
	fmt.Println(jellybeanCounter)

	var flavorScale float32 = 5.8
	fmt.Println(flavorScale)

	var nameOfSong string = "The Final Countdown"
	fmt.Println("Gob's theme song is " + nameOfSong)

	var emptyInt int8
	var emptyFloat float32
	var emptyString string
	fmt.Println(emptyInt, emptyFloat, emptyString)

	daysOnVacation := 7
	var hoursInDay = 24
	fmt.Println("You have spent", daysOnVacation * hoursInDay, "hours on vacation.")

	var cupsOfCoffeeConsumed int
	cupsOfCoffeeConsumed= 10
	fmt.Println("I've had " + strconv.Itoa(cupsOfCoffeeConsumed) + " of coffee today!") 

	coolSneakers := 65.99
  	niceNecklace := 45.50
	var taxCalculation float64
	taxCalculation += coolSneakers
	taxCalculation += niceNecklace
	taxCalculation *= 0.08875
	fmt.Println("Purchase of", coolSneakers + niceNecklace, "with 8.875% sales tax", taxCalculation, "equal to", coolSneakers + niceNecklace + taxCalculation)

	var magicNum, powerLevel int32
	magicNum= 42
	powerLevel= 9001
	fmt.Println("magicNum is:", magicNum, "powerLevel is:", powerLevel)
	amount, unit := 10, "doll hairs"
	fmt.Println(amount, unit, ", that's expensive...")
}