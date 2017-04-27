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

func CalculateSeats(totalPopulation int, peoplePerSeat int, seatLimit int) int {
	seats := 0

	for population := totalPopulation; population > peoplePerSeat; population -= peoplePerSeat {
		seats++
		if seats == seatLimit {
			peoplePerSeat += 10000
		}
	}

	return seats
}
