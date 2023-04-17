package game_stats

import (
	"encoding/csv"
	"fmt"
	"os"
)

func CsvImport() {
	file, err := os.Open("./data/games.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Println(records)
}
