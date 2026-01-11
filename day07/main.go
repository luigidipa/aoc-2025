package main

import (
	"fmt"

	"github.com/luigidipa/aoc-2025/utils"
)

func main() {
	part1()
}

func part1() {
	diagr := makeDiagr()
	splits := 0
	advance(diagr, 0, findStartPt(diagr), &splits)
	fmt.Println(splits)
}

func advance(diagr [][]rune, line, pos int, splits *int) {
	// Out of bounds
	if pos < 0 || pos >= len(diagr[0]) || line == len(diagr)-1 {
		return
	}
	// Path was already traversed
	if diagr[line+1][pos] == '|' {
		return
	}
	// Continue down
	if diagr[line+1][pos] == '.' {
		diagr[line+1][pos] = '|'
		advance(diagr, line+1, pos, splits)
		return
	}
	// Split
	*splits++
	advance(diagr, line+1, pos-1, splits)
	advance(diagr, line+1, pos+1, splits)
}

func makeDiagr() [][]rune {
	lines := utils.ReadLines("./input.txt")
	diagr := make([][]rune, len(lines))
	for i, l := range lines {
		diagr[i] = []rune(l)
	}
	return diagr
}

func findStartPt(diagr [][]rune) int {
	// Find start point
	for i, r := range diagr[0] {
		if r == 'S' {
			return i
		}
	}
	return -1 // Unreachable
}
