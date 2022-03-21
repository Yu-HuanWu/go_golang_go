package main

import "fmt"

func addHundred(num int) {
  num += 100
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
}