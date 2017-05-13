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
				"Results don't match, expected: %d, received: %d",
				testCase.expectedSeats,
				result,
			)
		}
	}
}
