package main

import (
	"fmt"
	"sort"
)

func nonConstructibleChange(coins []int) int {
	// sort the input array in place
	sort.Ints(coins)

	// create a variable to keep track of the current change that can be created
	change := 0
	// iterate through the array
	for _, coin := range coins {
		// if the current coin is greater than the change + 1, then we know that the change + 1 is the smallest amount of change that cannot be created
		if coin > change+1 {
			// if the above condition is true, then we return the change + 1
			return change + 1
		}
		// if the above condition is not true, then we add the current coin to the change
		change += coin
	}
	return change + 1
}

func main() {
	coins := []int{1, 1, 1, 1, 1, 1, 1}
	fmt.Println(nonConstructibleChange(coins))
}
