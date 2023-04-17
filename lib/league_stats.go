package league_stats

import (
	"fmt"
	"os"

	"github.com/kniren/gota/dataframe"
)

func LoadGameTeamsData() dataframe.DataFrame {
	file, err := os.Open("./data/games_teams.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	games_teamsDF := dataframe.ReadCSV(file)

	return games_teamsDF
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
}
