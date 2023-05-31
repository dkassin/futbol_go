package lib

func CalculateWinningestCoach(gameTeams []GameTeamsData, gamesData []GamesData, year string) string {
	winsByTeamMap := make(map[string]int)
	gamesByheadcoachMap := make(map[string]int)
	coachWinningPercentageMap := make(map[string]float64)
	var largestValueCoach	string
	var largestValue 		float64

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
		if len(largestValueCoach) == 0 || percentage > largestValue {
			largestValue = percentage
			largestValueCoach = coach
		}
	}
	
	return largestValueCoach
}