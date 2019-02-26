package strategy

import (
	"github.com/incazteca/congress_allocation/state"
)

// CalculateSeats calculate the seats for with the new strategy
func CalculateSeats(states []state.State) strategyResult {
	return newStrategySummary(states)
}

func newStrategySummary(states []state.State) strategySummary {
	var totalPopulation int
	var totalHouseSeats int

	for _, state := range states {
		totalPopulation += state.PopulationEst
	}
	return strategySummary{}
}
