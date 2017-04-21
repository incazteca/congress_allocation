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

	fmt.Println(states[0].Name)
}
