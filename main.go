package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
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
const currentHouseSeatTotal = 435
const senateSeats = 2
const censusYear = 2010

func main() {
	states := getStates()

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

func getStates() []State {
	fh, err := os.Open(sourceFile)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(fh)
	i := 0
	var states []State

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if i == 0 || record[0] == "—" {
			i++
			continue
		}

		states = append(states, newState(record))
	}

	return states
}

func newState(record []string) State {
	pop2017 := formattedToInt(record[3])
	pop2010 := formattedToInt(record[4])
	currentSeats := formattedToInt(record[6])

	return State{
		Name:             record[2],
		PopulationCensus: pop2010,
		PopulationEst:    pop2017,
		HouseSeats:       currentSeats,
		SenateSeats:      senateSeats,
		ElectoralVotes:   currentSeats + senateSeats,
		CensusYear:       censusYear,
	}
}

func formattedToInt(number string) int {
	commaReplace := strings.NewReplacer(",", "")
	result, err := strconv.Atoi(commaReplace.Replace(number))

	if err != nil {
		log.Fatal(err)
	}
	return result
}

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
