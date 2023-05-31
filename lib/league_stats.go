package lib

import (
	"encoding/csv"
	"math"
	"os"
	"strconv"
)

type GamesTeamsDataRaw struct {
	Data [][]string
	Headers map[string]int
}

type GameTeamsData struct {
	GameId		string
	TeamId		int
	HomeAway	string
	Result		string
	SettledIn	string
	HeadCoach	string
	Goals		int
	Shots		int
	Tackles		int
	Pim			int
	PowerPOpps	int
	FaceWin		float64
	GiveAways	int
	TakeAways	int
}

type TeamsDataRaw struct {
	Data [][]string
	Headers map[string]int
}

type TeamsData struct {
	TeamId		int
	FranchiseId	int
	TeamName	string
	Abrev		string
	Stadium		string
	Link		string
}



func LoadGameTeamsData() (GameTeamsDataRaw, error) {
	file, err := os.Open("./data/games_teams.csv")
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
	var gameTeamsData []gameTeamsData

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

		teamId, _ := strconv.Atoi(row[teamIdIndex])
		goals, _ := strconv.Atoi(row[goalsIndex])
		shots, _ := strconv.Atoi(row[shotsIndex])
		tackles, _ := strconv.Atoi(row[tacklesIndex])
		pim, _ := strconv.Atoi(row[pimIndex])
		powerPOpps, _ := strconv.Atoi(row[powerPlayOpportunitiesIndex])
		faceWin, _ := strconv.ParseFloat(row[faceOffWinPercentageIndex], 64)
		giveAways, _ := strconv.Atoi(row[giveawaysIndex])
		takeAways, _ := strconv.Atoi(row[takeawaysIndex])

		gameTeamData := gameTeamsData {
			GameId: 	row[gameIdIndex],
			TeamId:		teamId,
			HomeAway:	row[homeAwayIndex],
			Result:		row[resultIndex],
			SettledIn:	row[settledInIndex],
			HeadCoach:	row[headCoachIndex],
			Goals:		goals,
			Shots:		shots,
			Tackles:	tackles,
			Pim:		pim,
			powerPOpps	powerPOpps,
			FaceWin,	faceWin,
			GiveAways,	giveAways,
			TakeAways	takeAways,
		}

		gameTeamsData.append(gameTeamsData, gameTeamData)	
	}

	return gameTeamsData

}

func StructureTeamsData(teamsDataRaw TeamsDataRaw) []TeamsData {
	var teamsData []teamsData

	teamIdIndex := TeamsDataRaw.Headers["team_id"]
	franchiseIdIndex := TeamsDataRaw.Headers["franchiseId"]
	teamName := TeamsDataRaw.Headers["teamName"]
	abbrevIndex := TeamsDataRaw.Headers["abbreviation"]
	stadiumIndex := TeamsDataRaw.Headers["Stadium"]
	linkIndex := TeamsDataRaw.Headers["link"]

	for i, row := range teamsDataRaw.Data {
		if i == 0 {
			continue
		}

		teamId, _ := strconv.Atoi(row[teamIdIndex])
		franchiseId, _ := strconv.Atoi(row[franchiseIdIndex])
	

		teamData := teamsData {
			TeamId:			teamId,
			FranchiseId:	franchiseId,
			TeamName		row[teamName],
			Abrev			row[abbrevIndex],
			Stadium			row[stadiumIndex],
			Link			row[linkIndex],
		}

		teamsData.append(teamsData, teamData)
	}

	return teamsData
}
