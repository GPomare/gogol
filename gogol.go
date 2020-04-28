package main

import (
	"flag"
	"log"
)

func main() {
	fileInput := flag.String("input", "", "Text file that contains the input for Conway's Game of Life")
	iterations := flag.Int("iterations", 12, "The number of generations to simulate")
	flag.Parse()

	if *fileInput == "" {
		flag.PrintDefaults()
		return
	}

	data, err := LoadEnvironment(*fileInput)

	if err != nil {
		log.Fatal(err.Error())
	}

	err = ValidateInput(data)

	if err != nil {
		log.Fatal(err.Error())
	}

	var e Environment
	e.SetBoard(data)
	GameOfLife(&e, *iterations, true)
}
