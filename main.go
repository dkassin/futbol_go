package main

import (
	"fmt"
	"github.com/dkassin/futbol_go/lib"
)

func main() {
	df1_test, err := lib.LoadGameTeamsData()
	if err != nil {
		fmt.Println("Error loading game teams data:", err)
		return
	}

	lib.PrintData(df1_test)
}
