package execution

import (
	"testing"
)

func TestWashProfileData(t *testing.T) {

	err := Wash("profiledata.json", "washed.json")

	if err != nil {
		t.Errorf("Error washing data")
	}

}
