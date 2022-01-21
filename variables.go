package main
import (
	"fmt"
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
}