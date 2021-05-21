package main

import (

  "fmt"
)

func bubbleSort(arr [10]int ) []int {
  for i := 1 ; i < 10 ; i++ {
    for j := 0 ; j < 9 ; j++ {
      if arr[j] > arr[j+1] {
        fmt.Println("reached")
        Swap(arr[0:], j)
        
      } 
    }
  } 
return arr[0:]
}

func Swap(arr []int, i int)  {
    fmt.Println(arr[i])
        temp := (arr)[i]
        (arr)[i] = (arr)[i+1]
        (arr)[i+1] = temp
      
  
}

func main() {
  var arr [10]int ; var sorted []int
 

  
  fmt.Println("Enter")
  for i := 0 ; i < 10 ; i++ {
    fmt.Scan(&arr[i])
  }
  sorted =  bubbleSort(arr)
 
  fmt.Println(sorted)
}