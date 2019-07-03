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

func TestTotalSeats(t *testing.T) {
	testCases := []totalSeatTestCase{
		totalSeatTestCase{13, 100, 5, 5, 5},
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
		calculateSeatTestCase{9, 80, 10, 10},
		calculateSeatTestCase{9, 80, 10, 10},
		calculateSeatTestCase{0, 0, 10, 10},
		calculateSeatTestCase{1, 8, 10, 10},
		calculateSeatTestCase{6, 100, 20, 10},
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

func TestSeatsToAllocate(t *testing.T) {
	// test input and expected result
	testCases := [][]int{
		{0, 0},
		{100, 1},
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

		result := seatsToAllocate(population)

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
