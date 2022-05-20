package main

import "fmt"

func main() {
	bubbleArray := []int{1, 3, 9, -5, 0}

	for i := 0; i < len(bubbleArray); i++ {
		fmt.Println(bubbleArray[i])
	}
}
