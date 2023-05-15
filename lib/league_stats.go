package lib

import (
	"fmt"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func LoadGameTeamsData() (dataframe.DataFrame, error) {
	file, err := os.Open("./data/game_teams.csv")
	if err != nil {
		return dataframe.DataFrame{}, err
	}
	defer file.Close()

	games_teamsDF := dataframe.ReadCSV(file)

	return games_teamsDF, nil
}

func LoadTeamsData() dataframe.DataFrame {
	file2, err := os.Open("./data/teams.csv")
	if err != nil {
		panic(err)
	}
	defer file2.Close()

	teamsDF := dataframe.ReadCSV(file2)

	return teamsDF
}

func PrintData(df dataframe.DataFrame) {
	fmt.Println(df)
	fmt.Sprintf("%T", df)
}
