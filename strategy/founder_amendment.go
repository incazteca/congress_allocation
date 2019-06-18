package strategy

import (
	"github.com/incazteca/congress_allocation/state"
)

const strategyName = "Founder Amendment"
const senateSeatsPerState = 2

const initialPeoplePerSeat = 40000
const seatsPerStep = 100
const additionalPeoplePerSeatPerStep = 10000

// CalculateSeats calculate the seats for with the new strategy
// based on estimated population
func Allocate(states []state.State) strategyResult {
	orderedStates := state.SortOnEstimatedPopulation(states)
	totalPop := totalEstimatedPopulation(states)
	availableSeats := seatsToAllocate(totalPop)
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

func seatsToAllocate(population int) int {
	return TotalSeats(
		population, initialPeoplePerSeat, seatsPerStep, additionalPeoplePerSeatPerStep,
	)
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

// TotalSeats Get the total seats
func TotalSeats(population int, peoplePerSeat int, stepSeatLimit int, seatStep int) int {
	seats := 0
	workingSeats := 0

	for population >= peoplePerSeat {
		workingSeats = CalculateSeats(population, peoplePerSeat, stepSeatLimit)

		population -= (peoplePerSeat * workingSeats)
		peoplePerSeat += seatStep
		seats += workingSeats
		workingSeats = 0

		if stepSeatLimit == 0 {
			break
		}
	}

	return seats
}

// CalculateSeats Calculate seats by population
func CalculateSeats(population int, peoplePerSeat int, seatLimit int) int {
	seats := 0

	for population >= peoplePerSeat {
		seats++

		if seats >= seatLimit {
			break
		}

		population -= peoplePerSeat
	}

	return seats
}
