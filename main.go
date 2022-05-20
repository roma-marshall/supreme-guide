package main

import (
	"fmt"
)

func main() {
	bubbleArray := []int{1, 3, 9, -5, 0, 2, 77, 11}
	min := 0
	max := 0
	temp := 0
	for i := len(bubbleArray) - 1; i > 0; i-- {
		if bubbleArray[i] < min {
			min = bubbleArray[i]
		}
		if bubbleArray[i] > max {
			max = bubbleArray[i]
		}
		for j := 0; j < i; j++ {
			if bubbleArray[j] > bubbleArray[j+1] {
				temp = bubbleArray[j]
				bubbleArray[j] = bubbleArray[j+1]
				bubbleArray[j+1] = temp
			}
		}
	}
	for i := 0; i < len(bubbleArray); i++ {
		fmt.Print(bubbleArray[i], " ")
	}

	fmt.Println("\nmax: ", max)
	fmt.Println("min: ", min)
}

func Max(bubbleArray []int) {
	panic("unimplemented")
}
