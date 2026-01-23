package main

import (
	"fmt"

	"github.com/luigidipa/aoc-2025/utils"
)

func main() {
	part1()
	part2()
}

func part1() {
	diagr := makeDiagr()
	splits := advance(diagr, 0, findStartPt(diagr))
	fmt.Println(splits)
}

func advance(diagr [][]rune, r, c int) int {
	// Out of bounds
	if c < 0 || c >= len(diagr[0]) || r == len(diagr)-1 {
		return 0
	}
	// Path was already traversed
	if diagr[r+1][c] == '|' {
		return 0
	}
	// Continue down
	if diagr[r+1][c] == '.' {
		diagr[r+1][c] = '|'
		return advance(diagr, r+1, c)
	}
	// Split
	return 1 + advance(diagr, r+1, c-1) + advance(diagr, r+1, c+1)
}

func part2() {
	diagr := makeDiagr()
	timelnsCount := make([][]int, len(diagr))
	res := 0

	// Create matrix to keep count of timelines
	for i := range timelnsCount {
		timelnsCount[i] = make([]int, len(diagr[0]))
	}

	// Set point below startpoint to 1
	timelnsCount[1][findStartPt(diagr)] = 1

	for r := range len(diagr) - 1 {
		for c := range len(diagr[0]) {
			if timelnsCount[r][c] == 0 {
				continue
			}
			if diagr[r+1][c] == '.' {
				timelnsCount[r+1][c] += timelnsCount[r][c]
			} else {
				timelnsCount[r+1][c-1] += timelnsCount[r][c]
				timelnsCount[r+1][c+1] += timelnsCount[r][c]
			}
		}
	}

	// Res is the sum of the last row
	for _, t := range timelnsCount[len(timelnsCount)-1] {
		res += t
	}
	fmt.Println(res)
}

func makeDiagr() [][]rune {
	rs := utils.ReadLines("./input.txt")
	diagr := make([][]rune, len(rs))
	for i, l := range rs {
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
