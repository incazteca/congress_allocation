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
	}
}
