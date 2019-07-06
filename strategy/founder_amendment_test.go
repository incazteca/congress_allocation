package strategy

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

func TestSeatsAvailableToAllocate(t *testing.T) {
	// Each case is input, followed by expected output
	testCases := [][]int{
		{0, 0},
		{1, 1},
		{39999, 1},
		{40000, 1},
		{40001, 2},
		{4000000, 100},
		{4000001, 101},
		{4040000, 101},
		{4040001, 101},
		{4050000, 101},
		{4050001, 102},
	}

	for _, testCase := range testCases {
		population := testCase[0]
		expectation := testCase[1]

		result := SeatsAvailableToAllocate(population)

		if expectation != result {
			t.Errorf(
				"Results don't match, expected: %d, received %d, with input population of %d",
				expectation,
				result,
				population,
			)
		}
	}
}
