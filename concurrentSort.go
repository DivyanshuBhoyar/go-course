package main

import (
	"fmt"
	"sync"
)

func bubbleSort(arr []int, c chan []int, wg *sync.WaitGroup) {

	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				swap(arr[0:], j)

			}
		}
	}
	temp := arr[0:]
	c <- append(temp)
	fmt.Println(temp)
	wg.Done()

}

func swap(arr []int, i int) {
	temp := (arr)[i]
	(arr)[i] = (arr)[i+1]
	(arr)[i+1] = temp

}

func main() {
	var arr []int
	var sliceFinal []int
	sorted := make(chan []int, 5)
	var wg sync.WaitGroup
	var size, i, input int

	fmt.Println("Enter Size")
	fmt.Scanf("%d\n", &size)
	midIndex := size / 2

	fmt.Println("Enter individual elements")
	for i < size {
		fmt.Scanf("%d\n", &input)
		arr = append(arr, input)
		i++
	}
	wg.Add(4)
	go bubbleSort(arr[0:midIndex/2], sorted, &wg)
	go bubbleSort(arr[(midIndex/2):midIndex], sorted, &wg)
	go bubbleSort(arr[midIndex:(midIndex/2)+midIndex], sorted, &wg)
	go bubbleSort(arr[(midIndex/2)+midIndex:], sorted, &wg)
	wg.Wait()
	close(sorted)

	for i := range sorted {
		// fmt.Println(i)
		sliceFinal = append(sliceFinal, i...)
	}

	sliceFinal = func(arr []int) []int {
		for i := 1; i < len(arr); i++ {
			for j := 0; j < len(arr)-1; j++ {
				if arr[j] > arr[j+1] {
					swap(arr[0:], j)

				}
			}
		}
		return arr[0:]
	}(sliceFinal)

	fmt.Println("The sorted array is \n ", sliceFinal)

}
