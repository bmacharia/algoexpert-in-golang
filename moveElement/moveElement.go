// move element to the end of the array
package main

import (
	"fmt"
)

func moveElement(arr []int, element int) []int {
	// intialize the left and right pointers
	i := 0
	j := len(arr) - 1
	// while loop i less than j
	for i < j {
		// while i is less than j and arr[j] is equal to element
		for i < j && arr[j] == element {
			// decrement j
			j--
		}
		// if arr[i] is equal to element
		if arr[i] == element {
			// swap arr[i] and arr[j]
			arr[i], arr[j] = arr[j], arr[i]
		}
		// increment i
		i++
	}
	// return arr
	return arr
}

func main() {

	arr := []int{2, 1, 2, 2, 2, 3, 4, 2}
	fmt.Println(moveElement(arr, 2))
}
