package main

import (
	"fmt"
	"sort"
)

func threeNumberSum(array []int, target int) [][]int {
	// sort the input array
	sort.Ints(array)
	// create a 2 dimemsional array to store the result
	result := [][]int{}
	// iterate through the array up unitl second last element
	for i := 0; i < len(array)-2; i++ {
		// create two pointers, one at the start and one at the end
		left := i + 1
		right := len(array) - 1
		// iterate through the array until the left pointer is greater than the right pointer
		for left < right {
			// calculate the sum of the three numbers
			currentSum := array[i] + array[left] + array[right]
			//
			if currentSum == target {
				result = append(result, []int{array[i], array[left], array[right]})
				// move the left pointer to the right
				left++
				// move the right pointer to the left
				right--
			} else if currentSum < target {
				// move the left pointer to the right
				left++
			} else {
				// move the right pointer to the left
				right--
			}
		}
	}
	return result
}

func main() {
	array := []int{12, 3, 1, 2, -6, 5, -8, 6}
	target := 0
	result := threeNumberSum(array, target)
	fmt.Println(result)
}
