package main_test

import (
	"testing"

	"github.com/dkassin/futbol_go/lib"
)

func TestLoadGamesData(t *testing.T) {
	_, err := lib.LoadGamesData()
	if err != nil {
		t.Fatalf("Failed to load games data: %v", err)
	}
}

func TestLoadGameTeamsData(t *testing.T) {
	_, err := lib.LoadGameTeamsData()
	if err != nil {
		t.Errorf("Failed to load games data: %s\n err:%+v", err.Error(), err)
	}
}

func TestLoadTeamsData(t *testing.T) {
	_, err := lib.LoadGameTeamsData()
	if err != nil {
		t.Errorf("Failed to load games data: %s\n err:%+v", err.Error(), err)
	}
}