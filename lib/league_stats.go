package lib

import (
	"encoding/csv"
	// "math"
	"os"
	"strconv"
)

type GameTeamsDataRaw struct {
	Data [][]string
	Headers map[string]int
}

type GameTeamsData struct {
	GameId		string
	TeamId		string
	HomeAway	string
	Result		string
	SettledIn	string
	HeadCoach	string
	Goals		int
	Shots		int
	Tackles		int
	Pim			int
	PowerPOpps	int
	PowerPGoals int
	FaceWin		float64
	GiveAways	int
	TakeAways	int
}

type TeamsDataRaw struct {
	Data [][]string
	Headers map[string]int
}

type TeamsData struct {
	TeamId		string
	FranchiseId	string
	TeamName	string
	Abrev		string
	Stadium		string
	Link		string
}



func LoadGameTeamsData() (GameTeamsDataRaw, error) {
	file, err := os.Open("./data/game_teams.csv")
	if err != nil {
		return GameTeamsDataRaw{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return GameTeamsDataRaw{}, err
	}

	headers := make(map[string]int)
	if len(data) > 1 {
		for i, header := range  data[0] {
			headers[header] = i
		}
	}

	return GameTeamsDataRaw{
		Data: data,
		Headers: headers,
		}, nil
}


func LoadTeamsData() (TeamsDataRaw, error) {
	file, err := os.Open("./data/teams.csv")
	if err != nil {
		return TeamsDataRaw{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return TeamsDataRaw{}, err
	}

	headers := make(map[string]int)
	if len(data) > 1 {
		for i, header := range  data[0] {
			headers[header] = i
		}
	}

	return TeamsDataRaw{
		Data: data,
		Headers: headers,
		}, nil
}

func StructureGameTeamsData(gameTeamsDataRaw  GameTeamsDataRaw) []GameTeamsData {
	var gameTeams []GameTeamsData

	gameIdIndex := gameTeamsDataRaw.Headers["game_id"]
	teamIdIndex := gameTeamsDataRaw.Headers["team_id"]
	homeAwayIndex := gameTeamsDataRaw.Headers["HoA"]
	resultIndex := gameTeamsDataRaw.Headers["result"]
	settledInIndex := gameTeamsDataRaw.Headers["settled_in"]
	headCoachIndex := gameTeamsDataRaw.Headers["head_coach"]
	goalsIndex := gameTeamsDataRaw.Headers["goals"]
	shotsIndex := gameTeamsDataRaw.Headers["shots"]
	tacklesIndex := gameTeamsDataRaw.Headers["tackles"]
	pimIndex := gameTeamsDataRaw.Headers["pim"]
	powerPlayOpportunitiesIndex:= gameTeamsDataRaw.Headers["powerPlayOpportunities"]
	powerPlayGoalsIndex := gameTeamsDataRaw.Headers["powerPlayGoals"]
	faceOffWinPercentageIndex := gameTeamsDataRaw.Headers["faceOffWinPercentage"]
	giveawaysIndex := gameTeamsDataRaw.Headers["giveaways"]
	takeawaysIndex := gameTeamsDataRaw.Headers["takeaways"]

	for i, row := range gameTeamsDataRaw.Data {
		if i == 0 {
			continue
		}

		goals, _ := strconv.Atoi(row[goalsIndex])
		shots, _ := strconv.Atoi(row[shotsIndex])
		tackles, _ := strconv.Atoi(row[tacklesIndex])
		pim, _ := strconv.Atoi(row[pimIndex])
		powerPOpps, _ := strconv.Atoi(row[powerPlayOpportunitiesIndex])
		powerPGoals, _ := strconv.Atoi(row[powerPlayGoalsIndex])
		faceWin, _ := strconv.ParseFloat(row[faceOffWinPercentageIndex], 64)
		giveAways, _ := strconv.Atoi(row[giveawaysIndex])
		takeAways, _ := strconv.Atoi(row[takeawaysIndex])

		gameTeam := GameTeamsData {
			GameId: 		row[gameIdIndex],
			TeamId:			row[teamIdIndex],
			HomeAway:		row[homeAwayIndex],
			Result:			row[resultIndex],
			SettledIn:		row[settledInIndex],
			HeadCoach:		row[headCoachIndex],
			Goals:			goals,
			Shots:			shots,
			Tackles:		tackles,
			Pim:			pim,
			PowerPOpps:		powerPOpps,
			PowerPGoals: 	powerPGoals,
			FaceWin:		faceWin,
			GiveAways:		giveAways,
			TakeAways:		takeAways,
		}

		gameTeams = append(gameTeams, gameTeam)	
	}

	return gameTeams

}

func StructureTeamsData(teamsDataRaw TeamsDataRaw) []TeamsData {
	var teams []TeamsData

	teamIdIndex := teamsDataRaw.Headers["team_id"]
	franchiseIdIndex := teamsDataRaw.Headers["franchiseId"]
	teamName := teamsDataRaw.Headers["teamName"]
	abbrevIndex := teamsDataRaw.Headers["abbreviation"]
	stadiumIndex := teamsDataRaw.Headers["Stadium"]
	linkIndex := teamsDataRaw.Headers["link"]

	for i, row := range teamsDataRaw.Data {
		if i == 0 {
			continue
		}
	
		teamData := TeamsData {
			TeamId:			row[teamIdIndex],
			FranchiseId:	row[franchiseIdIndex],
			TeamName:		row[teamName],
			Abrev:			row[abbrevIndex],
			Stadium:		row[stadiumIndex],
			Link:			row[linkIndex],
		}

		teams = append(teams, teamData)
	}

	return teams
}

func CalculateCountOfTeams(teams []TeamsData) int {
	return len(teams)
}

func CalculateBestOffense(gameTeams []GameTeamsData) string {
	goalsByTeamMap := make(map[string]int)
	gamesByTeamMap := make(map[string]int)
	var largestValueTeam 	string
	var largestValue 		float64

	for _, gameTeam := range gameTeams {
		goalsByTeamMap[gameTeam.TeamId] += gameTeam.Goals	
		gamesByTeamMap[gameTeam.TeamId]++ 
	}

	for team, games := range gamesByTeamMap {
		avgGoals := float64(goalsByTeamMap[team]) / float64(games)
		if largestValueTeam == "" || largestValue < avgGoals {
			largestValue = avgGoals
			largestValueTeam = team
		} 			
	}

	return largestValueTeam
}

func ReturnTeamName(teams []TeamsData, largestValueTeam string) string {
	for _, team := range teams {
		if team.TeamId == largestValueTeam {
			return team.TeamName
		}
	} 
	return ""
}