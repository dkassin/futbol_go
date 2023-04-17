package league_stats_test

import (
	"futbol_go/lib/league_stats"
	"testing"
)

func TestLoadGameTeamsData(t *testing.T) {
	df := league_stats.LoadGameTeamsData()
	if _, ok := df.(*dataframe.DataFrame); !ok {
		t.Errorf("Returned value is not a dataframe")
	}
}
