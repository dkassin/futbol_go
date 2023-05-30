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
	percentageTies := lib.CalculatePercentageTies(structuredData)
	gameCountMapBySeason := lib.CalculateCountGamesBySeason(structuredData)
	avgGoalsPerGame := lib.CalculateAverageGoalsPerGame(structuredData)
	avgGoalsPerSeason := lib.CalculateAverageGoalsBySeason(structuredData, gameCountMapBySeason)
	fmt.Println("highest Total Score:", highestTotalScore)
	fmt.Println("lowest Total Score:", lowestTotalScore)
	fmt.Println("Percentage Home Wins:", percentageHomewins)
	fmt.Println("Percentage Ties:", percentageTies)
	fmt.Println("Game Counts by Season:", gameCountMapBySeason)
	fmt.Println("Avg Goals per Game:", avgGoalsPerGame)
	fmt.Println("Avg Goals per each Season:", avgGoalsPerSeason)
}
