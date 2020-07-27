
/**
 * Ron Nathaniel
 * July 27 2020
 */

package main

import (
  "fmt"
  "math"
)


func IsPythTriplet(triplet ...int) bool {
  return math.Pow(float64(triplet[0]), 2) +
    math.Pow(float64(triplet[1]), 2) == math.Pow(float64(triplet[2]), 2)
}

func IsSumReached(triplet ...int) bool {
  sum := 0
  for i := range triplet {
    sum += triplet[i]
  }

  return sum == 1000
}


func GetTriangle() (triangle [3]int) {
  for i := 2; i < 1000; i++ {
    for j := 1; j < i; j++ {
      for k := 0; k < j; k++ {
        if IsPythTriplet(k, j, i) && IsSumReached(k, j, i) {
          triangle = [3]int{k, j, i}
          return
        }
      }
    }
  }
  return
}


func main() {
  t := GetTriangle()
  fmt.Println(t)
}
