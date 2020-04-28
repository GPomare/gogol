package main

import (
	"fmt"
	"testing"
)

func TestLoadEnvironment(t *testing.T) {
	env, err := LoadEnvironment("./data/repeator.txt")
	if err != nil {
		t.Error(err.Error())
	} else if len(env) != 5 {
		t.Error("Failed to load environment")
		fmt.Print(env)
	}
}

func TestLoadEnvironmentInvalidPath(t *testing.T) {
	_, err := LoadEnvironment("./data/doesnotexist.txt")
	if err == nil {
		t.Error("Failed to produce error")
	}
}

func TestValidInputInvalidLength(t *testing.T) {
	err := ValidateInput([][]string{{"-", "-", "-"}, {"-", "-"}, {"-", "-", "-"}})
	if err == nil {
		t.Error("Did not throw error with invalid row length")
	}
}

func TestValidInputInvalidCharacter(t *testing.T) {
	err := ValidateInput([][]string{{"-", "S", "-"}, {"-", "-", "-"}})
	if err == nil {
		t.Error("Did not throw error with invalid character")
	}
}

func TestValidInput(t *testing.T) {
	err := ValidateInput([][]string{{"-", "x", "-"}, {"-", "-", "-"}})
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGameOfLife(t *testing.T) {
	data, _ := LoadEnvironment("./data/repeator.txt")
	var env Environment

	testCases := []int{1, 99, 1001}

	for _, val := range testCases {
		env.SetBoard(data)
		GameOfLife(&env, val, false)
		if env.Cells[2][1] != "x" || env.Cells[2][3] != "x" {
			t.Error("cells not set correctly")
			env.PrintCells()
			fmt.Println()
		}
	}
}
