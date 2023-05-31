package lib

import (
	"encoding/csv"
	"math"
	"os"
	"strconv"
)

type GamesDataRaw struct {
	Data [][]string
	Headers map[string]int
}

type GamesData struct {
	Date		string
	HomeTeamId	int
	AwayTeamId	int
	HomeGoals	int
	AwayGoals	int
	GameId		string
	Season		string
	Type		string
	Venue		string
}

func LoadGamesData() (GamesDataRaw, error) {
	file, err := os.Open("./data/games.csv")
	if err != nil {
		return GamesDataRaw{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return GamesDataRaw{}, err
	}

	headers := make(map[string]int)
	if len(data) > 1 {
		for i, header := range  data[0] {
			headers[header] = i
		}
	}

	return GamesDataRaw{
		Data: data,
		Headers: headers,
		}, nil
}

func StructureGamesData(gamesDataRaw GamesDataRaw) []GamesData {
	var gamesData []GamesData

	dateTimeIndex := gamesDataRaw.Headers["date_time"]
	homeTeamIdIndex := gamesDataRaw.Headers["home_team_id"]
	awayTeamIdIndex := gamesDataRaw.Headers["away_team_id"]
	homeGoalsIndex := gamesDataRaw.Headers["home_goals"]
	awayGoalsIndex := gamesDataRaw.Headers["away_goals"]
	gameIdIndex := gamesDataRaw.Headers["game_id"]
	seasonIndex := gamesDataRaw.Headers["season"]
	typeIndex := gamesDataRaw.Headers["type"]
	venueIndex := gamesDataRaw.Headers["venue"]

	for i, row := range gamesDataRaw.Data {
		if i == 0 {
			continue
		}

		homeTeamId, _ := strconv.Atoi(row[homeTeamIdIndex])
		awayTeamId, _ := strconv.Atoi(row[awayTeamIdIndex])
		homeGoals, _ := strconv.Atoi(row[homeGoalsIndex])
		awayGoals, _ := strconv.Atoi(row[awayGoalsIndex])
	
		gameData := GamesData{
			Date:		row[dateTimeIndex],
			HomeTeamId:	homeTeamId,
			AwayTeamId:	awayTeamId,
			HomeGoals:	homeGoals,
			AwayGoals:	awayGoals,
			GameId:		row[gameIdIndex],
			Season:		row[seasonIndex],
			Type:		row[typeIndex],
			Venue:		row[venueIndex],
		}

		gamesData = append(gamesData, gameData)
	}	

	return gamesData
}

func CalculateHighestTotalScore(gamesData []GamesData) int {
	highestScore := 0

	for _, gameData := range gamesData {
		totalScore := gameData.HomeGoals + gameData.AwayGoals

		if totalScore > highestScore {
			highestScore = totalScore
		}
	}
	return highestScore
}

func CalculateLowestTotalScore(gamesData []GamesData) int {
	lowestScore := 0

	for _, gameData := range gamesData {
		totalScore := gameData.HomeGoals + gameData.AwayGoals
		if totalScore < lowestScore {
			lowestScore = totalScore
		}
	}
	return lowestScore
}

func CalculatePercentageHomeWins(gamesData []GamesData) float64 {
	homeWins := 0
	totalGames := 0

	for _, gameData := range gamesData {
		if gameData.HomeGoals > gameData.AwayGoals {
			homeWins++
			totalGames++
		} else {
			totalGames++
		}
	}
	winPercentage := float64(homeWins) / float64(totalGames)
	roundedPercentage := math.Round(winPercentage*100) / 100
	return roundedPercentage
}

func CalculatePercentageTies(gamesData []GamesData) float64 {
	ties := 0
	totalGames := 0

	for _, gameData := range gamesData {
		if gameData.HomeGoals == gameData.AwayGoals {
			ties++
			totalGames++
		} else {
			totalGames++
		}
	}
	tiePercentage := float64(ties) / float64(totalGames)
	roundedPercentage := math.Round(tiePercentage*100) / 100
	return roundedPercentage
}

func CalculateCountGamesBySeason(gamesData []GamesData) map[string]int {
	gameCountMap := make(map[string]int)

	for _, gameData := range gamesData {
		if _, exists := gameCountMap[gameData.Season]; exists {
			gameCountMap[gameData.Season]++
		} else {
			gameCountMap[gameData.Season] = 1
		}
	}
	return gameCountMap
}

func CalculateAverageGoalsPerGame(gamesData []GamesData) float64 {
	totalGoals := 0
	totalGames := 0

	for _, gameData := range gamesData {
		totalGoals += gameData.HomeGoals + gameData.AwayGoals
		totalGames++
	}
	avgGoals := float64(totalGoals) / float64(totalGames)
	roundedPercentage := math.Round(avgGoals*100) / 100
	return roundedPercentage
}

func CalculateAverageGoalsBySeason(gamesData []GamesData, gameCountMap map[string]int) map[string]float64{
	goalsBySeasonMap := make(map[string]float64)


	for _, gameData := range gamesData {
		goalsBySeasonMap[gameData.Season] += float64(gameData.HomeGoals + gameData.AwayGoals)
	}

	for year, goals := range goalsBySeasonMap {
		avgGoalsPerYear := goals / float64(gameCountMap[year])
		 goalsBySeasonMap[year] = math.Round(avgGoalsPerYear*100) / 100
	}
	return goalsBySeasonMap
}