package main

import (
	"testing"
)

type calculateSeatTestCase struct {
	expectedSeats int
	Population    int
	PeoplePerSeat int
	SeatLimit     int
}

type totalSeatTestCase struct {
	expectedSeats int
	Population    int
	PeoplePerSeat int
	StepSeatLimit int
	SeatStep      int
}

type formattedToIntTestCase struct {
	input    string
	expected int
}

func TestTotalSeats(t *testing.T) {
	testCases := []totalSeatTestCase{
		totalSeatTestCase{11, 100, 5, 5, 5},
		totalSeatTestCase{20, 100, 5, 5, 0},
		totalSeatTestCase{20, 100, 5, 20, 0},
		totalSeatTestCase{20, 100, 5, 20, 20},
		totalSeatTestCase{6, 100, 5, 3, 20},
		totalSeatTestCase{29, 100, 1, 5, 1},
	}

	for _, testCase := range testCases {
		result := TotalSeats(
			testCase.Population,
			testCase.PeoplePerSeat,
			testCase.StepSeatLimit,
			testCase.SeatStep,
		)

		if testCase.expectedSeats != result {
			t.Errorf(
				"Results for total don't match, expected: %d, received: %d With population of %d, seatlimit of %d, people per seat %d, and seat step of %d",
				testCase.expectedSeats,
				result,
				testCase.Population,
				testCase.StepSeatLimit,
				testCase.PeoplePerSeat,
				testCase.SeatStep,
			)
		}
	}
}

func TestCalculateSeats(t *testing.T) {
	testCases := []calculateSeatTestCase{
		calculateSeatTestCase{10, 100, 10, 10},
		calculateSeatTestCase{10, 1000, 10, 10},
		calculateSeatTestCase{8, 80, 10, 10},
		calculateSeatTestCase{8, 80, 10, 10},
		calculateSeatTestCase{0, 0, 10, 10},
		calculateSeatTestCase{5, 100, 20, 10},
		calculateSeatTestCase{3, 100, 20, 3},
		calculateSeatTestCase{3, 100, 20, 3},
	}

	for _, testCase := range testCases {
		result := CalculateSeats(
			testCase.Population,
			testCase.PeoplePerSeat,
			testCase.SeatLimit,
		)

		if testCase.expectedSeats != result {
			t.Errorf(
				"Results don't match, expected: %d, received: %d With population of %d, seatlimit of %d and people per seat %d",
				testCase.expectedSeats,
				result,
				testCase.Population,
				testCase.SeatLimit,
				testCase.PeoplePerSeat,
			)
		}
	}
}

func TestformattedToInt(t *testing.T) {
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

func TestnewState(t *testing.T) {
	record := []string{"", "", "Chicago", "2,700,000", "2,750,000", "-", "3"}
	expectedState := State{
		Name:             "Chicago",
		PopulationCensus: 2700000,
		PopulationEst:    2750000,
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
