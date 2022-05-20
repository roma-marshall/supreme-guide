package main

import "fmt"

func main() {
	bubbleArray := []int{1, 3, 9, -5, 0, 2, 77, 11}
	max := 0
	min := 0
	for i := 0; i < len(bubbleArray); i++ {
		fmt.Print(bubbleArray[i], " ")
		if bubbleArray[i] > max {
			max = bubbleArray[i]
		}
		if bubbleArray[i] < min {
			min = bubbleArray[i]
		}
	}
	fmt.Println("\nmax: ", max)
	fmt.Println("min: ", min)
}
