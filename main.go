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
	gameTeamsDataRaw, err := lib.LoadGameTeamsData()
	if err != nil {
		panic(err)
	}
	teamsDataRaw, err := lib.LoadTeamsData()
	if err != nil {
		panic(err)
	}

	structuredGamesData := lib.StructureGamesData(gamesDataRaw)
	structuredGameTeamsData := lib.StructureGameTeamsData(gameTeamsDataRaw)
	structuredTeamsData := lib.StructureTeamsData(teamsDataRaw)
	highestTotalScore := lib.CalculateHighestTotalScore(structuredGamesData)
	lowestTotalScore := lib.CalculateLowestTotalScore(structuredGamesData)
	percentageHomewins := lib.CalculatePercentageHomeWins(structuredGamesData)
	percentageTies := lib.CalculatePercentageTies(structuredGamesData)
	gameCountMapBySeason := lib.CalculateCountGamesBySeason(structuredGamesData)
	avgGoalsPerGame := lib.CalculateAverageGoalsPerGame(structuredGamesData)
	avgGoalsPerSeason := lib.CalculateAverageGoalsBySeason(structuredGamesData, gameCountMapBySeason)
	countOfTeams := lib.CalculateCountOfTeams(structuredTeamsData)
	bestOffense := lib.ReturnTeamName(structuredTeamsData, lib.CalculateBestOffense(structuredGameTeamsData))
	fmt.Println("highest Total Score:", highestTotalScore)
	fmt.Println("lowest Total Score:", lowestTotalScore)
	fmt.Println("Percentage Home Wins:", percentageHomewins)
	fmt.Println("Percentage Ties:", percentageTies)
	fmt.Println("Game Counts by Season:", gameCountMapBySeason)
	fmt.Println("Avg Goals per Game:", avgGoalsPerGame)
	fmt.Println("Avg Goals per each Season:", avgGoalsPerSeason)
	fmt.Println("Count of Teams:", countOfTeams)
	fmt.Println("Best Offense:", bestOffense)
}
