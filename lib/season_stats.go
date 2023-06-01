package lib

func CalculateBestWorstCoach(gameTeams []GameTeamsData, gamesData []GamesData, year string) (string, string, string) {
	winsByTeamMap := make(map[string]int)
	gamesByheadcoachMap := make(map[string]int)
	coachWinningPercentageMap := make(map[string]float64)
	var largestValueCoach, smallestValueCoach	string
	var largestValue, smallestValue  		float64	

	for _, game := range gamesData {
		if game.Season == year && game.Type == "Regular Season" {
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

