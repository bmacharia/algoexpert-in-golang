package twoSum

// brute force method using two for loops
//O(n^2) time | O(1) space
func TwoSum(nums []int, target int) []int {
	for i := 1; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{nums[i], nums[j]}
			}
		}

	}
	// return empty if no two vales sum up to the target value
	return []int{}

}

//
func TwoSumHash(array []int, target int) []int {

	// create a hash table to store values and for fast access
	numInSet := map[int]bool{}

	// iterate over the array one element at a time
	for _, num := range array {

		// each iteration calculate the value that we needed
		valNeeded := target - num

		//check the hash table if the value Needed exists in the set
		// if it exists return the values to the calling function
		if _, found := numInSet[valNeeded]; found {
			return []int{valNeeded, num}

		} // if the values does not exist insert the value into the hash table and set its value to true
		numInSet[num] = true

	} // if no two values add up to the target value return an empty slice to the caller
	return []int{}

}
