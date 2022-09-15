package main

import "fmt"

func isMonotonic(array []int) bool {
	isNonDecreasing := true
	isNonIncreasing := true

	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			isNonDecreasing = false
		}
		if array[i] > array[i-1] {
			isNonIncreasing = false
		}
	}
	return isNonDecreasing || isNonIncreasing
}

func main() {
	array := []int{-1, -5, -10, -1100, -1100, -1101, -1102, -9001}
	fmt.Println(isMonotonic(array))
}
