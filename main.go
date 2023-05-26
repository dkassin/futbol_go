package main

import (
	"fmt"

	"github.com/dkassin/futbol_go/lib"
)

func main() {
	gamesDataRaw, err := lib.LoadGamesData()
	if err != nil {
		panic(err)
	}

	structuredData := lib.StructureGamesData(gamesDataRaw)
	highestTotalScore := lib.CalculateHighestTotalScore(structuredData)
	lowestTotalScore := lib.CalculateLowestTotalScore(structuredData)
	percentageHomewins := lib.CalculatePercentageHomeWins(structuredData)
	fmt.Println("highest Total Score:", highestTotalScore)
	fmt.Println("lowest Total Score:", lowestTotalScore)
	fmt.Println("Percentage Home Wins:", percentageHomewins)
}
