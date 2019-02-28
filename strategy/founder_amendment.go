package strategy

import (
	"github.com/incazteca/congress_allocation/state"
	"sort"
)

const strategyName = "Founder Amendment"
const senateSeatsPerState = 2

const initialPeoplePerSeat = 40000
const seatStep = 100
const AdditionalPeoplePerSeat = 10000

// CalculateSeats calculate the seats for with the new strategy
// based on estimated population
func Allocate(states []state.State) strategyResult {
	summary := newStrategySummary(states)

	sort.Slice(states, func(i, j int) bool {
		return states[i].PopulationEst < states[j].PopulationEst
	})

	return strategyResult{
		Name:    strategyName,
		States:  states, // Actually work allocations here
		Summary: summary,
	}
}

func newStrategySummary(states []state.State) strategySummary {
	var totalPopulationEst int
	var totalHouseSeats int
	var totalSenateSeats int

	for _, state := range states {
		totalPopulationEst += state.PopulationEst
		totalSenateSeats += senateSeatsPerState
		totalHouseSeats += calculateHouseSeatsForState(state)
	}

	for _, state := range states {
		totalHouseSeats += calculateHouseSeatsForState(state)
	}

	return strategySummary{
		totalPopulation:  totalPopulationEst,
		totalHouseSeats:  totalHouseSeats,
		totalSenateSeats: totalSenateSeats,
	}
}

func calculateHouseSeatsForState(currentState state.State) int {
	return len(currentState.Name)
}
