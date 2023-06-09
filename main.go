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
	worstOffense := lib.ReturnTeamName(structuredTeamsData, lib.CalculateWorstOffense(structuredGameTeamsData))
	highestScoringVisitor := lib.ReturnTeamName(structuredTeamsData, lib.CalculateHighestScoringVisitor(structuredGameTeamsData))
	highestScoringHome := lib.ReturnTeamName(structuredTeamsData, lib.CalculateHighestScoringHome(structuredGameTeamsData))
	lowestScoringVisitor := lib.ReturnTeamName(structuredTeamsData, lib.CalculateLowestScoringVisitor(structuredGameTeamsData))
	lowestScoringHome := lib.ReturnTeamName(structuredTeamsData, lib.CalculateLowestScoringHome(structuredGameTeamsData))
	yearCoach, winningestCoach, worstCoach := lib.CalculateBestWorstCoach(structuredGameTeamsData, structuredGamesData, "20132014")
	yearShots, mostAccurateTeamId, leastAccurateTeamID := lib.CalculateBestWorstAccurateTeam(structuredGameTeamsData, structuredGamesData, "20132014")
	mostAccurateTeam := lib.ReturnTeamName(structuredTeamsData, mostAccurateTeamId)
	leastAccurateTeam := lib.ReturnTeamName(structuredTeamsData, leastAccurateTeamID)
	yearTackles, mostTacklesTeamId, leastTacklesTeamID := lib.CalculateBestWorstTacklesTeam(structuredGameTeamsData, structuredGamesData, "20132014")
	mostTacklesTeam := lib.ReturnTeamName(structuredTeamsData, mostTacklesTeamId)
	leastTacklesTeam := lib.ReturnTeamName(structuredTeamsData, leastTacklesTeamID)
	fmt.Println("highest Total Score:", highestTotalScore)
	fmt.Println("lowest Total Score:", lowestTotalScore)
	fmt.Println("Percentage Home Wins:", percentageHomewins)
	fmt.Println("Percentage Ties:", percentageTies)
	fmt.Println("Game Counts by Season:", gameCountMapBySeason)
	fmt.Println("Avg Goals per Game:", avgGoalsPerGame)
	fmt.Println("Avg Goals per each Season:", avgGoalsPerSeason)
	fmt.Println("Count of Teams:", countOfTeams)
	fmt.Println("Best Offense:", bestOffense)
	fmt.Println("Worst Offense:", worstOffense)
	fmt.Println("Highest Scoring Visitor:", highestScoringVisitor)
	fmt.Println("Highest Scoring Home Team:", highestScoringHome)
	fmt.Println("Lowest Scoring Visitor:", lowestScoringVisitor)
	fmt.Println("Lowest Scoring Home Team:", lowestScoringHome)
	fmt.Println(yearCoach,"Winningest Coach:", winningestCoach)
	fmt.Println(yearCoach,"Worst Coach:", worstCoach)
	fmt.Println(yearShots,"Most Accurate Team:", mostAccurateTeam)
	fmt.Println(yearShots,"Least Accurate Team:", leastAccurateTeam)
	fmt.Println(yearTackles,"Most Tackles Team:", mostTacklesTeam)
	fmt.Println(yearTackles,"Least Tackles Team:", leastTacklesTeam)
}
