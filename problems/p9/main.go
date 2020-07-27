
package main

import (
  "fmt"
  "math"
)


func IsPythTriplet(triplet ...int) bool {
  // fmt.Printf("%d, %d, %d\n", triplet[0], triplet[1], triplet[2])
  fmt.Println(triplet)
  return math.Pow(float64(triplet[0]), 2) +
    math.Pow(float64(triplet[1]), 2) == math.Pow(float64(triplet[2]), 2)
}


func GetTriangle() (triangle [3]int) {
  for i := 2; i < 100000; i++ {
    for j := 1; j < i; j++ {
      for k := 0; k < j; k++ {
        if IsPythTriplet(k, j, i) {
          fmt.Println("IS")
          // triangle = {k, j, i}
          triangle[0] = k
          triangle[1] = j
          triangle[2] = i
          break
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
