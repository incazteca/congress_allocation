package state

import (
	"testing"
)

type formattedToIntTestCase struct {
	input    string
	expected int
}

func TestNewState(t *testing.T) {
	record := []string{"", "", "Chicago", "2,700,000", "2,750,000", "-", "3"}
	expectedState := State{
		Name:             "Chicago",
		PopulationEst:    2700000,
		PopulationCensus: 2750000,
		HouseSeats:       3,
		SenateSeats:      2,
		ElectoralVotes:   5,
		CensusYear:       2010,
	}
	state := newState(record)

	if expectedState != state {
		t.Errorf("Expected %v, recieved %v", expectedState, state)
	}
}

func TestFormattedToInt(t *testing.T) {
	testCases := []formattedToIntTestCase{
		formattedToIntTestCase{"1", 1},
		formattedToIntTestCase{"12", 12},
		formattedToIntTestCase{"-12", -12},
		formattedToIntTestCase{"0", 0},
		formattedToIntTestCase{"1,000", 1000},
		formattedToIntTestCase{"234,234", 234234},
	}

	for _, testCase := range testCases {
		result := formattedToInt(testCase.input)
		if testCase.expected != result {
			t.Errorf("Expected %d, recieved: %d", testCase.expected, result)
		}
	}
}

func TestSortOnEstimatedPopulation(t *testing.T) {
	lowPopState := State{Name: "Low", PopulationEst: 30}
	medPopState := State{Name: "Med", PopulationEst: 50}
	highPopState := State{Name: "Med", PopulationEst: 100}

	unorderedStates := []State{medPopState, lowPopState, highPopState}
	expectedStates := []State{lowPopState, medPopState, highPopState}

	actual := SortOnEstimatedPopulation(unorderedStates)

	for i, expected := range expectedStates {
		if expected != actual[i] {
			t.Errorf("Expected +%v, received: +%v", expected, actual)
		}
	}
}
