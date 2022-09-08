package TournamentWinner

func TournamentWinner(competitions [][]string, results []int) string {

	// create a map to store the points of each team
	scores := map[string]int{}

	//the variable winner to store the winning team
	winner := ""

	//loop through the competitions array only using the index
	for i := range competitions {

		// store the results of each match from the competitions array in localWinner variable
		localWinner := competitions[i][1-results[i]]
		// store the points of the localWinner in the scores map
		scores[localWinner] += 3
		// check if the points of the localWinner is greater than the points of the winner
		if scores[localWinner] > scores[winner] {
			// if true, store the localWinner in the winner variable
			winner = localWinner
		}
	}
	// return the winner
	return winner
}
