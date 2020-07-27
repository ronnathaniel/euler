
/**
 * Ron Nathaniel
 * July 26 2020
 */
 
package main

import (
  "fmt"
)


func Total(limit int) (total int) {
  total = 0
  for i := 0; i < limit; i++ {
    if i % 3 == 0 || i % 5 == 0 {
      total += i
    }
  }
  return
}

func main() {
  t := Total(1000)
  fmt.Println(t)
}
