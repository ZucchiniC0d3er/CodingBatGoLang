package main

import "fmt"

func main() {
   fmt.Println(monkeyTrouble(true, false));
}
/**
*if both monkeys are smiling then we are in trouble
*if both are not smiling then we are also in trouble
**/
func monkeyTrouble(aSmile, bSmile bool) bool {

   if ((aSmile && bSmile) || (!aSmile && !bSmile)) {
      return true;
   }
   return false;
}