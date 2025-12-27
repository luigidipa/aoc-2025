package main

import (
	"fmt"
	"strings"

	"github.com/sssdrum/aoc-2025/utils"
)

func main() {
	part1()
	part2()
}

func part1() {
	res := 0
	grid := makeGrid("./input.txt")
	rows, cols := len(grid), len(grid[0])
	dirs := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}, {-1, 1}, {-1, -1}, {1, 1}, {1, -1}}
	for i := range rows {
		for j := range cols {
			if grid[i][j] != "@" {
				continue
			}

			neighs := 0
			for _, d := range dirs {
				ni, nj := i+d[0], j+d[1]
				if ni < 0 || ni >= rows || nj < 0 || nj >= cols {
					continue
				}
				if grid[ni][nj] == "@" {
					neighs++
				}
			}
			if neighs < 4 {
				res++
			}
		}
	}
	fmt.Println(res)
}

func part2() {
	res := 0
	grid := makeGrid("./input.txt")
	rows, cols := len(grid), len(grid[0])
	dirs := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}, {-1, 1}, {-1, -1}, {1, 1}, {1, -1}}

	for {
		removed := false
		for i := range rows {
			for j := range cols {
				if grid[i][j] != "@" {
					continue
				}

				neighs := 0
				for _, d := range dirs {
					ni, nj := i+d[0], j+d[1]
					if ni < 0 || ni >= rows || nj < 0 || nj >= cols {
						continue
					}
					if grid[ni][nj] == "@" {
						neighs++
					}
				}
				if neighs < 4 {
					grid[i][j] = "."
					res++
					removed = true
				}
			}
		}
		if !removed { // Break if no rolls were removed in this iteration
			break
		}
	}
	fmt.Println(res)
}

func makeGrid(path string) [][]string {
	lines := strings.Split(utils.ReadInput(path), "\n")

	// Initialize grid
	rows, cols := len(lines), len(lines[0])
	grid := make([][]string, rows)
	for i := range grid {
		grid[i] = make([]string, cols)
	}

	// Fill grid
	for i, line := range lines {
		for j, val := range line {
			grid[i][j] = string(val)
		}
	}

	return grid
}
