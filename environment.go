package main

import (
	"fmt"
	"strings"
)

// Environment tracks the current board state
type Environment struct {
	Cells         [][]string
	Width, Height int
}

// GetNumAliveNeighbours returns the number of alive neighbours around given cell
func (e *Environment) GetNumAliveNeighbours(pos [2]int) int {
	numAlive := 0
	row, col := pos[0], pos[1]

	var neighbours = [8][2]int{{row, col - 1}, {row - 1, col - 1}, {row - 1, col}, {row - 1, col + 1},
		{row, col + 1}, {row + 1, col + 1}, {row + 1, col}, {row + 1, col - 1}}

	for _, value := range neighbours {
		row := value[0]
		colID := value[1]

		if row < 0 {
			row = len(e.Cells) - 1
		} else if row >= len(e.Cells) {
			row = 0
		}

		if colID < 0 {
			colID = len(e.Cells[0]) - 1
		} else if colID >= len(e.Cells[0]) {
			colID = 0
		}

		if e.Cells[row][colID] == "x" {
			numAlive++
		}
	}

	return numAlive
}

// SetBoard updates the environment to a new state
func (e *Environment) SetBoard(board [][]string) {
	e.Cells = board
	e.Height = len(board)
	e.Width = len(board[0])
}

// PrintCells outputs the board state to the terminal
func (e *Environment) PrintCells() {
	fmt.Println()
	for _, value := range e.Cells {
		fmt.Println(strings.Join(value, ""))
	}
}
