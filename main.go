package main

import (
	"encoding/json"
	"fmt"
	"github.com/incazteca/congress_allocation/state"
	"io/ioutil"
	"log"
)

// State has the relevant data for a state
type State struct {
	Name             string `json:"name"`
	PopulationCensus int    `json:"population_census"`
	PopulationEst    int    `json:"population_estimate"`
	HouseSeats       int    `json:"house_seats"`
	SenateSeats      int    `json:"senate_seats"`
	ElectoralVotes   int    `json:"electoral_votes"`
	CensusYear       int    `json:"year"`
}

const sourceFile = "data/us_population.csv"

func main() {
	states := state.GetStates(sourceFile)

	var totalPopulation int
	var totalHouseSeats int

	for _, state := range states {
		totalPopulation += state.PopulationCensus
		totalHouseSeats += state.HouseSeats
	}

	newSeats := TotalSeats(totalPopulation, 40000, 100, 10000)

	fmt.Printf("Total population: %d \n", totalPopulation)
	fmt.Printf("Total house seats: %d \n", totalHouseSeats)
	fmt.Printf("Total house seats by new method: %d \n", newSeats)
}

func getStatesJSON() []State {
	fileName := "data/states.json"
	raw, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}
	var states []State
	json.Unmarshal(raw, &states)
	return states
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
