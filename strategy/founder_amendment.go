package strategy

import (
	"github.com/incazteca/congress_allocation/state"
)

const senateSeatsPerState = 2

// CalculateSeats calculate the seats for with the new strategy
// based on estimated population
func CalculateSeats(states []state.State) strategyResult {
	return newStrategySummary(states)
}

func newStrategySummary(states []state.State) strategySummary {
	var totalPopulationEst int
	var totalHouseSeats int

	for _, state := range states {
		totalPopulationEst += state.PopulationEst
		totalHouseSeats += calculateHouseSeatsForState(state)
		totalSenateSeats += senateSeatsPerState
	}

	return strategySummary{
		totalPopulation:  totalPopulationEst,
		totalHouseSeats:  totalHouseSeats,
		totalSenateSeats: totalSenateSeats,
	}
}

func calculateHouseSeatsForState(currentState state) {

}
