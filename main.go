package main

import (
	"fmt"
)

func main() {
	bubbleArray := []int{1, 3, 9, -5, 0, 2, 77, 11}

	fmt.Println(bubbleArray)
	fmt.Println(bubleSort(bubbleArray))
	fmt.Println("min: ", minValue(bubbleArray))
	fmt.Println("max: ", maxValue(bubbleArray))
}
func bubleSort(bubbleArray []int) []int {
	temp := 0
	for i := len(bubbleArray) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if bubbleArray[j] > bubbleArray[j+1] {
				temp = bubbleArray[j]
				bubbleArray[j] = bubbleArray[j+1]
				bubbleArray[j+1] = temp
			}
		}
	}
	return bubbleArray
}

func minValue(bubbleArray []int) int {
	min := 0
	for i := len(bubbleArray) - 1; i > 0; i-- {
		if bubbleArray[i] < min {
			min = bubbleArray[i]
		}
	}

	return min
}

func maxValue(bubbleArray []int) int {
	max := 0
	for i := len(bubbleArray) - 1; i > 0; i-- {
		if bubbleArray[i] < max {
			max = bubbleArray[i]
		}
	}

	return max
}
