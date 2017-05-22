package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type State struct {
	Name       string         `json:"name"`
	Population int            `json:"population"`
	Seats      map[string]int `json:"seats"`
}

func main() {
	fileName := "data/states.json"
	fmt.Println("vim-go")
	raw, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var states []State
	json.Unmarshal(raw, &states)

	var totalPopulation int
	var totalHouseSeats int

	for _, state := range states {
		totalPopulation += state.Population
		totalHouseSeats += state.Seats["house"]
	}

	newSeats := CalculateSeats(totalPopulation, 40000, 100)

	fmt.Printf("Total population: %d \n", totalPopulation)
	fmt.Printf("Total house seats: %d \n", totalHouseSeats)
	fmt.Printf("Total house seats by new method: %d \n", newSeats)
}

func TotalSeats(population int, peoplePerSeat int, stepSeatLimit int, seatStep int) int {
	seats := 0
	workingSeats := 0

	for working_population := population; working_population >= peoplePerSeat; working_population -= (peoplePerSeat * workingSeats) {
		workingSeats = CalculateSeats(population, peoplePerSeat, stepSeatLimit)
		peoplePerSeat += seatStep
		seats += workingSeats
	}

	return seats
}

func CalculateSeats(population int, peoplePerSeat int, seatLimit int) int {
	seats := 0

	for working_population := population; working_population >= peoplePerSeat; working_population -= peoplePerSeat {
		seats++

		if seats >= seatLimit {
			break
		}
	}

	return seats
}
