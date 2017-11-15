package main

import "fmt"
import "math"

func main() {
   fmt.Println(diff21(21));
}
/**
*returns the absolute value between n and 21
*if n is greater then 21 the double the absolute value
**/
func diff21(n float64) float64 {

   if(n > 21){
       return (math.Abs(21 - n) *2);
   }
   return (21 - n);
}