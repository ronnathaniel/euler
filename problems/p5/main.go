
package main

import (
  "fmt"
)


func SmallestDivisible () (smallest int) {
  NextNum:
    for smallest = 1; true; smallest++ {
      fmt.Println(smallest)
      for i := 1; i < 20; i++ {
        if smallest % i != 0 {
          continue NextNum
        }
      }
      break
    }
    return
}

func main() {
  t := SmallestDivisible()
  fmt.Println(t)

}
