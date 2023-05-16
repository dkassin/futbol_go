package main

import (
	"fmt"

	"github.com/dkassin/futbol_go/lib"
)

func main() {
	gamesData, err := lib.LoadGamesData()
	if err != nil {
		panic(err)
	}

	highestTotalScore := lib.CalculateHighestTotalScore(gamesData)
	lowestTotalScore := lib.CalculateLowestTotalScore(gamesData)
	percentageHomewins := lib.CalculatePercentageHomeWins(gamesData)
	fmt.Println("highest Total Score:", highestTotalScore)
	fmt.Println("lowest Total Score:", lowestTotalScore)
	fmt.Println("Percentage Home Wins:", percentageHomewins)
}
