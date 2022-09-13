package main

import (
	"fmt"
	"math"
	"sort"
)

func smallesDifference(arrayOne, arrayTwo []int) []int {
	// sort the input arrays
	sort.Ints(arrayOne)
	sort.Ints(arrayTwo)
	//intialize two pointers
	idxOne, idxTwo := 0, 0
	// initialize the smallest difference with the max value
	smallest := math.MaxInt32
	// initialize the current difference with the max value
	current := math.MaxInt32
	// create a slice to store the result
	smallestPair := []int{}
	// loop through the arrays
	for idxOne < len(arrayOne) && idxTwo < len(arrayTwo) {
		// get the first element from the first array and the second element from the second array
		firstNum, secondNum := arrayOne[idxOne], arrayTwo[idxTwo]
		// if firtNum is less than secondNum
		if firstNum < secondNum {
			// set the current to the difference between the secondNum and the firstNum
			current = secondNum - firstNum
			// increment the idxOne
			idxOne++
			// else if the secondNum is less than the firstNum
		} else if secondNum < firstNum {
			// set the current to the difference between the firstNum and the secondNum
			current = firstNum - secondNum
			//
			idxTwo++
		} else {
			// return the pair if firstNum and secondNum are equal
			return []int{firstNum, secondNum}
		}
		// if the current is less than the smallest
		if smallest > current {
			// set the smallest to the current
			smallest = current
			// set the smallestPair to the pair of firstNum and secondNum
			smallestPair = []int{firstNum, secondNum}
		}
	}
	// return the smallestPair
	return smallestPair
}

func main() {
	arrayOne := []int{-1, 5, 10, 20, 28, 3}
	arrayTwo := []int{26, 134, 135, 15, 17}
	result := smallesDifference(arrayOne, arrayTwo)
	fmt.Println(result)

}
