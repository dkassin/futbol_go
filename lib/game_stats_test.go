package lib_test

import (
	"testing"

	"github.com/dkassin/futbol_go/lib"
)

func TestCalculateHighestTotalScoreEmptyStruct(t *testing.T) {
	score := lib.CalculateHighestTotalScore([]lib.GamesData{})
	if score != 0 {
		t.Errorf("CalculateHighestTotalScore([]lib.GamesData{}) = %d; got 0", score)
	}
}

func TestCalculateHighestTotalScore(t *testing.T) {
	value = []lib.GamesData{}
	score := lib.CalculateHighestTotalScore([]lib.GamesData{})
	if score != 0 {
		t.Errorf("CalculateHighestTotalScore([]lib.GamesData{}) = %d; got 0", score)
	}
}
