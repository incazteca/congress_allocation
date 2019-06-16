package state

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const currentHouseSeatTotal = 435
const senateSeats = 2
const censusYear = 2010

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

// GetStates gets states from a CSV file
func GetStates(sourceFile string) []State {
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

func sortOnEstimatedPopulation(states []State) []State {
	copyStates := states[:]
	sort.Slice(copyStates, func(i, j int) bool {
		return copyStates[i].PopulationEst < copyStates[j].PopulationEst
	})

	return copyStates
}

func formattedToInt(number string) int {
	commaReplace := strings.NewReplacer(",", "")
	result, err := strconv.Atoi(commaReplace.Replace(number))

	if err != nil {
		log.Fatal(err)
	}
	return result
}
