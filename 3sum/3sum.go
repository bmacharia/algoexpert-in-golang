package main

import (
	"fmt"
	"sort"
)

func threeNumberSum(array []int, target int) [][]int {
	// sort the input
	sort.Ints(array)
	// allocate a slice of slices for the output
	result := [][]int{}
	// iterate over the input up until the second to last element
	for i := 0; i < len(array)-2; i++ {
		// set a left pointer to the next element 
		left := i + 1
		// set a right pointer to the last element
		right := len(array) - 1
		// iterate while the left pointer is less than the right pointer
		for left < right {
			// calculate the sum of the three numbers
			currentSum := array[i] + array[left] + array[right]
			// if the sum is equal to the target
			if currentSum == target {
				//
				result = append(result, []int{array[i], array[left], array[right]})
				// increment the left pointer
				left++
				// decrement the right pointer
				right--
				// otherwise if the sum is less than the target
			} else if currentSum < target {
				// increment the left pointer
				left++
			} else {
				// decrement the right pointer
				right--
			}
		}
	}
	// return the output
	return result
}

func main() {
	array := []int{12, 3, 1, 2, -6, 5, -8, 6}
	target := 0
	result := threeNumberSum(array, target)
	fmt.Println(result)
}
