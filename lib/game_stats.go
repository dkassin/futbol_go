package lib

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
)

type GamesData struct {
	Data [][]string
}

func LoadGamesData() (GamesData, error) {
	file, err := os.Open("./data/games.csv")
	if err != nil {
		return GamesData{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return GamesData{}, err
	}

	return GamesData{Data: data}, nil
}

func CalculateHighestTotalScore(gamesData GamesData) int {

	highestScore := 0

	homeGoalsIndex := 6
	awayGoalsIndex := 7

	for i, row := range gamesData.Data {

		if i == 0 {
			continue
		}

		homeGoalsStr := row[homeGoalsIndex]
		awayGoalsStr := row[awayGoalsIndex]

		homeGoals, err := strconv.Atoi(homeGoalsStr)
		if err != nil {
			fmt.Printf("Conversion error: %v\n", err)
		}

		awayGoals, err := strconv.Atoi(awayGoalsStr)
		if err != nil {
			fmt.Printf("Conversion error: %v\n", err)
		}

		totalScore := homeGoals + awayGoals

		if totalScore > highestScore {
			highestScore = totalScore
		}
	}
	return highestScore
}

func CalculateLowestTotalScore(gamesData GamesData) int {
	lowestScore := 0

	homeGoalsIndex := 6
	awayGoalsIndex := 7

	for i, row := range gamesData.Data {

		if i == 0 {
			continue
		}

		homeGoalsStr := row[homeGoalsIndex]
		awayGoalsStr := row[awayGoalsIndex]

		homeGoals, err := strconv.Atoi(homeGoalsStr)
		if err != nil {
			fmt.Printf("Conversion error: %v\n", err)
		}

		awayGoals, err := strconv.Atoi(awayGoalsStr)
		if err != nil {
			fmt.Printf("Conversion error: %v\n", err)
		}

		totalScore := homeGoals + awayGoals

		if totalScore < lowestScore {
			lowestScore = totalScore
		}
	}
	return lowestScore
}

func CalculatePercentageHomeWins(gamesData GamesData) float64 {
	homeWins := 0
	totalGames := 0

	homeGoalsIndex := 6
	awayGoalsIndex := 7

	for i, row := range gamesData.Data {
		if i == 0 {
			continue
		}

		homeGoalsStr := row[homeGoalsIndex]
		awayGoalsStr := row[awayGoalsIndex]

		homeGoals, err := strconv.Atoi(homeGoalsStr)
		if err != nil {
			fmt.Printf("Conversion error: %v\n", err)
		}

		awayGoals, err := strconv.Atoi(awayGoalsStr)
		if err != nil {
			fmt.Printf("Conversion error: %v\n", err)
		}

		if homeGoals > awayGoals {
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
