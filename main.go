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

	fmt.Printf("Total population: %d \n", totalPopulation)
	fmt.Printf("Total house seats: %d \n", totalHouseSeats)
}

func CalculateSeatTotal(totalPopulation int, peoplePerSeat int, seatLimit int) {

}
