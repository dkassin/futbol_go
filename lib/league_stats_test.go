package lib_test

import (
	"testing"

	"github.com/dkassin/futbol_go/lib"
	// "github.com/go-gota/gota/dataframe"
)

func TestLoadGameTeamsData(t *testing.T) {
	_, err := lib.LoadGameTeamsData()
	if err != nil {
		t.Errorf("Returned value is not a dataframe, errstr: %s\n err:%+v", err.Error(), err)
	}
}
