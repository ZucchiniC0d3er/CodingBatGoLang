package main

import "fmt"
import "math"

func main() {
   fmt.Println(nearHundred(93));
}
/**
*returns true if the absolute value of the number that is 
*passed in is with in 10 of 100 or 200
**/
func nearHundred(n float64) bool {

  if ((math.Abs(100 - n) <= 10) || (math.Abs(200 - n) <= 10)){
      return true;
  }
  return false;
}