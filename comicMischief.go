package main
import (
	"fmt"
)

func main(){
  var publisher, writer, artist, title string
  var year, pageNumber uint
  var grade float32
  fmt.Println(title, "written by", writer, "drawn by", artist)
  title = "Mr. GoToSleep"
  writer= "Tracey Hatchet"
  artist= "Jewel Tampson"
  publisher= "DizzyBooks Publishing Inc."
  year= 1997
  pageNumber= 14
  grade= 6.5
  fmt.Println(title, "written by", writer, "drawn by", artist)

  title= "Epic Vol. 1"
  writer= "Ryan N. Shawn"
  artist= "Phoebe Paperclips"
  year= 2013
  pageNumber= 160
  grade= 9.0

  fmt.Println(title, "written by", writer, "drawn by", artist, "condition is", grade, "Number of pages:", pageNumber, "came out in", year, "published by", publisher)
}