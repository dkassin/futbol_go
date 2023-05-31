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
	takeaways	int
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
}