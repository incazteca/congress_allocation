package strategy

import (
	"github.com/incazteca/congress_allocation/state"
)

const strategyName = "Founder Amendment"
const senateSeatsPerState = 2

const initialPeoplePerSeat = 40000
const seatsPerStep = 100
const additionalPeoplePerSeatPerStep = 10000

// Allocate calculates the seats for the new strategy and allocates them to the
// states provided
func Allocate(states []state.State) strategyResult {
	orderedStates := state.SortOnEstimatedPopulation(states)
	totalPop := totalEstimatedPopulation(states)
	availableSeats := SeatsAvailableToAllocate(totalPop)
	allocatedStates := allocatePerState(availableSeats, orderedStates)

	return strategyResult{
		Name:    strategyName,
		States:  allocatedStates, // Actually work allocations here
		Summary: newStrategySummary(allocatedStates),
	}
}

func totalEstimatedPopulation(states []state.State) int {
	var population int

	for _, st := range states {
		population += st.PopulationEst
	}

	return population
}

// SeatsAvailableToAllocate Calculates seats available to allocate based on population
func SeatsAvailableToAllocate(population int) int {
	if population == 0 {
		return 0
	}

	step := 0
	seats := 0
	workingPop := population
	var popPerSeat int

	for workingPop > 0 {
		if seats != 0 && seats%100 == 0 {
			step++
		}
		popPerSeat = initialPeoplePerSeat + (additionalPeoplePerSeatPerStep * step)
		workingPop -= popPerSeat
		seats++
	}

	return seats
}

func allocatePerState(seatsAvailable int, states []state.State) []state.State {
	var allocatedStates []state.State

	return allocatedStates
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
