package main
 
import "fmt"

func eatTacos() {
  fmt.Println("Yum!")
}

func startGame() {
  instructions := "Press enter to start..." 
  fmt.Println(instructions)
}

func main() {
  eatTacos()
  startGame()
}