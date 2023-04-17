package main

import (
	"github.com/dkassin/futbol_go/league_stats"
)

func main() {
	df1_test := league_stats.LoadGameTeamsData()
	league_stats.PrintData(df1_test)
}
