package main

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strings"
)

// LoadEnvironment returns a stromg slice from file path
func LoadEnvironment(path string) ([][]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, strings.Split(line, ""))
	}

	return lines, nil
}

// ValidateInput determines whether input data is in a correct format.
// File must only contain "-" or "X" characters.
// All rows must be of identical length.
func ValidateInput(data [][]string) error {
	length := len(data[0])
	re := regexp.MustCompile("^[-x]*$")

	for _, row := range data {
		if len(row) != length {
			return errors.New("data contains rows of differing length")
		}
		for _, cell := range row {
			if re.MatchString(cell) == false {
				return errors.New("cell contains invalid character " + cell)
			}
		}
	}

	return nil
}

// GameOfLife performs gogol simulation for given iterations
func GameOfLife(env *Environment, iterations int, verbose bool) {
	if verbose {
		env.PrintCells()
	}

	for itNum := 0; itNum < iterations; itNum++ {
		newBoard := [][]string{}
		for rowID, rowVal := range env.Cells {
			newRow := []string{}
			for colID, colVal := range rowVal {
				numAlive := env.GetNumAliveNeighbours([2]int{rowID, colID})
				if colVal == "x" {
					if numAlive < 2 || numAlive > 3 {
						newRow = append(newRow, "-")
					} else {
						newRow = append(newRow, "x")
					}
				} else {
					if numAlive == 3 {
						newRow = append(newRow, "x")
					} else {
						newRow = append(newRow, "-")
					}
				}
			}
			newBoard = append(newBoard, newRow)
		}
		env.SetBoard(newBoard)
		if verbose {
			env.PrintCells()
		}
	}
}
