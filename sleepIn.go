package main

import "fmt"

func main() {
   fmt.Println(sleepIn(false, false));
}

/* function returning the max between two numbers */
func sleepIn(weekday, vaction bool) bool {

   if (!weekday || vaction) {
      return true;
   }
   return false;
}