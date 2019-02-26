package strategy

import (
	"github.com/incazteca/congress_allocation/state"
)

type strategy interface {
	CalculateSeats(states []state.State) strategyResult
}

type strategyResult struct {
	Name    string
	States  []state.State
	Summary strategySummary
}

// Let's have a summary of the modified states
type strategySummary struct {
	totalPopulation  int
	totalHouseSeats  int
	totalSenateSeats int
}
