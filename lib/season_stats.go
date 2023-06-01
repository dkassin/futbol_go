package lib

import (
	"math"
)
	
func CalculateBestWorstCoach(gameTeams []GameTeamsData, gamesData []GamesData, year string) (string, string, string) {
	winsByTeamMap := make(map[string]int)
	gamesByheadcoachMap := make(map[string]int)
	coachWinningPercentageMap := make(map[string]float64)
	var largestValueCoach, smallestValueCoach	string
	var largestValue, smallestValue  		float64	

	for _, game := range gamesData {
		if game.Season == year  {
			for _, gameTeam := range gameTeams {
				if game.GameId == gameTeam.GameId {
					if gameTeam.Result == "WIN" {
						winsByTeamMap[gameTeam.HeadCoach]++	
						gamesByheadcoachMap[gameTeam.HeadCoach]++ 
					} else {
						gamesByheadcoachMap[gameTeam.HeadCoach]++ 
					}
					
				}
			}
		}
	}

	for coach, games := range gamesByheadcoachMap {
		winPercentage := float64(winsByTeamMap[coach]) / float64(games)		
		coachWinningPercentageMap[coach] = winPercentage
	}

	for coach, percentage := range coachWinningPercentageMap {
		if len(largestValueCoach) == 0 {
			largestValue = percentage
			smallestValue = percentage
			largestValueCoach = coach
			smallestValueCoach = coach
		} else if percentage > largestValue {
			largestValue = percentage
			largestValueCoach = coach
		} else if percentage < smallestValue {
			smallestValue = percentage
			smallestValueCoach = coach
		}
	}

	return year, largestValueCoach, smallestValueCoach
}


func CalculateBestWorstAccurateTeam(gameTeams []GameTeamsData, gamesData []GamesData, year string) (string, string, string) {
	shotsByTeam := make(map[string]float64)
	goalsByTeam := make(map[string]float64)
	PercentageMap := make(map[string]float64)
	var largestValueRatio, smallestValueRatio	string
	var largestValue, smallestValue  		float64	

	for _, game := range gamesData {
		if game.Season == year {
			for _, gameTeam := range gameTeams {
				if game.GameId == gameTeam.GameId {
					shotsByTeam[gameTeam.TeamId] += float64(gameTeam.Shots)
					goalsByTeam[gameTeam.TeamId] += float64(gameTeam.Goals)
				}
			}
		}
	}

	for team, shots := range shotsByTeam {
		floatShots := shots
		floatGoals := goalsByTeam[team]
		goalPercentage := floatGoals /floatShots
		PercentageMap[team] = math.Round(goalPercentage*100) / 100
	}

	for team, percentage := range PercentageMap {
		if len(largestValueRatio) == 0 {
			largestValue = percentage
			smallestValue = percentage
			largestValueRatio = team
			smallestValueRatio = team
		} else if percentage > largestValue {
			largestValue = percentage
			largestValueRatio = team
		} else if percentage < smallestValue {
			smallestValue = percentage
			smallestValueRatio = team
		}
	}

	return year, largestValueRatio, smallestValueRatio
}
func CalculateBestWorstTacklesTeam(gameTeams []GameTeamsData, gamesData []GamesData, year string) (string, string, string) {
	tacklesByTeam := make(map[string]int)
	var largestValueTeam, smallestValueTeam	string
	var  largestValue, smallestValue  		int	

	for _, game := range gamesData {
		if game.Season == year  {
			for _, gameTeam := range gameTeams {
				if game.GameId == gameTeam.GameId {
					tacklesByTeam[gameTeam.TeamId] += gameTeam.Tackles
				}
			}
		}
	}

	for team, tackles := range tacklesByTeam {
		if len(largestValueTeam) == 0 {
			largestValue = tackles
			smallestValue = tackles
			largestValueTeam = team
			smallestValueTeam = team
		} else if tackles > largestValue {
			largestValue = tackles
			largestValueTeam = team
		} else if tackles < smallestValue {
			smallestValue = tackles
			smallestValueTeam = team
		}
	}

	return year, largestValueTeam, smallestValueTeam
}