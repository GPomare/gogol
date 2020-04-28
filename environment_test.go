package main

import "testing"

var testCases = []struct {
	cells    [][]string
	position [2]int
	expected int
}{
	{
		[][]string{{"0", "x", "0"}, {"0", "x", "0"}, {"0", "0", "x"}},
		[2]int{1, 1},
		2,
	},
	{
		[][]string{{"0", "0", "0"}, {"0", "0", "0"}, {"0", "0", "x"}},
		[2]int{1, 1},
		1,
	},
	{
		[][]string{{"0", "0", "0"}, {"0", "0", "0"}, {"0", "0", "0"}},
		[2]int{1, 1},
		0,
	},
	{
		[][]string{{"x", "x", "x"}, {"x", "x", "x"}, {"x", "x", "x"}},
		[2]int{1, 2},
		8,
	},
	{
		[][]string{{"x", "x", "x"}, {"x", "x", "x"}, {"x", "x", "x"}},
		[2]int{0, 0},
		8,
	},
	{
		[][]string{{"x", "x", "x"}, {"x", "x", "x"}, {"x", "x", "x"}},
		[2]int{2, 2},
		8,
	},
	{
		[][]string{{"x", "x", "x"}, {"x", "x", "x"}, {"x", "x", "x"}},
		[2]int{2, 0},
		8,
	},
	{
		[][]string{{"x", "x", "x"}, {"x", "x", "x"}, {"x", "x", "x"}},
		[2]int{0, 2},
		8,
	},
}

func TestGetNumAliveNeighbours(t *testing.T) {
	var e Environment
	for _, tests := range testCases {
		e.Cells = tests.cells
		result := e.GetNumAliveNeighbours(tests.position)
		if result != tests.expected {
			t.Error("Expected:", tests.expected, "Returned:", result)
		}
	}
}

func TestInitialize(t *testing.T) {
	var e Environment
	e.SetBoard([][]string{{"0", "0", "0"}})
	if e.Width != 3 || e.Height != 1 || e.Cells[0][0] != "0" {
		t.Error("Failed to initialize board")
	}
}
