package main

import (
  "fmt"
  "math"
)

func GenDisplaceFn(a, v0, s0 float64) func(t float64) float64 {

 return func (t float64) float64 {
    return v0 * t + 0.5 * a * math.Pow(t,2) - s0
  }

}

func main() {
  var acc , initVelo , initDisplacement  float64
 
  fmt.Println("Enter acc, init velocity & displacement all SEPERATED by a SINGLE SPACE")
  fmt.Scanf("%f %f %f\n" , &acc , &initVelo, &initDisplacement )

  fn := GenDisplaceFn(acc , initVelo , initDisplacement)

  var time float64

  fmt.Println("Enter time")
  fmt.Scanf("%f", &time)

  fmt.Println("the distance travlled is", fn(time))
}