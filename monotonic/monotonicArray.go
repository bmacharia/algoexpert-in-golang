package main

import "fmt"

func isMonotonic(array []int) bool {
	// non-decreasing means the integer values in the array are increasing
	isNonDecreasing := true
	// non-increasing means the integer values are in the array are decreasing
	isNonIncreasing := true
	// iterate through the array starting from the second element
	for i := 1; i < len(array); i++ {
		// if the current element is less than the previous element, the interger values are trending downwards
		if array[i] < array[i-1] {
			isNonDecreasing = false
		}
		// if the current element is greater than the previous element, the interger values are trending upwards
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
