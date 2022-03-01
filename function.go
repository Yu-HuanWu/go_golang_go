package main
 
import (
	"fmt"
	"time"
	"math"
)

func eatTacos() {
  fmt.Println("Yum!")
}

func startGame() {
  instructions := "Press enter to start..." 
  fmt.Println(instructions)
}

func isItLateInNewYork() string {
  var lateMessage string
  t := time.Now()
  tz, _ := time.LoadLocation("America/New_York")
  nyHour := t.In(tz).Hour()
  if nyHour < 5 {
    lateMessage = "Goodness it is late"
  } else if nyHour < 16 {
    lateMessage = "It's not late at all!"
  } else if nyHour < 19 {
    lateMessage = "I guess it's getting kind of late"
  } else {
    lateMessage = "It's late"
  }
  return lateMessage
}

func computeMarsYears(earthYears int) int {
  earthDays := earthYears * 365
  marsYears := earthDays / 687
  return marsYears
}

func specialComputation(x float64) float64 {
  return math.Log2(math.Sqrt(math.Tan(x)))
}

func combinedSpecialComputation(a, b, c, d float64) (float64, float64, float64, float64) {
	return specialComputation(a), specialComputation(b), specialComputation(c), specialComputation(d)
}

func getLikesAndShares(postId int) (int, int){
  var likesForPost, sharesForPost int
  switch postId {
  case 1:
    likesForPost = 5
    sharesForPost = 7
  case 2:
    likesForPost = 3
    sharesForPost = 11
  case 3:
    likesForPost = 22
    sharesForPost = 1
  case 4:
    likesForPost = 7
    sharesForPost = 9
  }
  fmt.Println("Likes: ", likesForPost, "Shares: ", sharesForPost)
  return likesForPost, sharesForPost
}

func main() {
  eatTacos()
  startGame()
  fmt.Println(isItLateInNewYork())
  fmt.Println("On Mars I am", computeMarsYears(33), "years old!")
  fmt.Println(combinedSpecialComputation(.0214, 1.02, 0.312, 4.001))
	
  var likes, shares int
  likes, shares = getLikesAndShares(4)
  if likes > 5 {
    fmt.Println("Woohoo! We got some likes.")
  }
  if shares > 10 {
    fmt.Println("We went viral!")
  }
}