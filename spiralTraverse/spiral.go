package main

import "fmt"

func spiralTraverse(array [][]int) []int {
	// check the length of the array if it is 0 return empty array
	if len(array) == 0 {
		return []int{}
	}
	// create a result array
	spiral := []int{}
	// create a top row pointer
	top_row := 0
	// create a bottom row pointer
	bottom_row := len(array) - 1
	// create a left column pointer
	left_column := 0
	// create a right column pointer
	right_column := len(array[0]) - 1

	// loop through the 2d array
	for top_row <= bottom_row && left_column <= right_column {
		//iterate top row of matrix
		for i := left_column; i <= right_column; i++ {
			// append top row elements to outout array
			spiral = append(spiral, array[top_row][i])

		}
		// iterate the last column of the matrix
		for i := top_row + 1; i <= bottom_row; i++ {
			// append last column elements to output array
			spiral = append(spiral, array[i][right_column])

		}
		// iterate the bottom row of the matrix
		for i := right_column - 1; i >= left_column; i-- {
			// check if the top row is equal to the bottom row
			if top_row == bottom_row {
				break
			}
			// append the bottom row elements to the output array
			spiral = append(spiral, array[bottom_row][i])

		}
		// iterate the first column of the matrix
		for i := bottom_row - 1; i > top_row; i-- {
			//check if the left column is equal to the right column
			if left_column == right_column {
				break
			}
			// append the first column elements to the output array
			spiral = append(spiral, array[i][left_column])

		}
		// increment the top row pointer
		top_row++
		// decrement the bottom row pointer
		bottom_row--
		// increment the left column pointer
		left_column++
		// decrement the right column pointer
		right_column--

	}
	return spiral
}

func main() {
	array := [][]int{
		{1, 2, 3, 4},
		{12, 13, 14, 5},
		{11, 16, 15, 6},
		{10, 9, 8, 7},
	}
	spiral := spiralTraverse(array)
	fmt.Println(spiral)
}
