package validSubsequence

// O(n) time and O(1) space
func IsValidSubsequence(array []int, sequence []int) bool {

	arrIdx := 0
	seqIdx := 0

	// while this condition is true
	for arrIdx < len(array) && seqIdx < len(sequence) {

		// check to see if the element in the array is equal to the element in the sequence
		// if equal incrent seqIdx variable, increment arrIdx varaible and repeat process for the next element
		if array[arrIdx] == sequence[seqIdx] {
			seqIdx += 1
		}
		arrIdx += 1
	}

	// this expression will evaluate to either true or false indicating a valid or invalid subsequence
	return seqIdx == len(sequence)
}
