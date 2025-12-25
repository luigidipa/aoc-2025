package main

import (
	"fmt"
	"strconv"

	"github.com/sssdrum/aoc-2025/utils"
)

func main() {
	part1()
}

func part1() {
	res := 0
	lines := utils.ReadLines("./input.txt")
	for _, line := range lines {
		currMax := 0
		l := 0
		for r := l + 1; r < len(line); r++ {
			a, _ := strconv.Atoi(string(line[l]))
			b, _ := strconv.Atoi(string(line[r]))
			currMax = max(currMax, a*10+b)
			if b > a {
				l = r
			}
		}
		res += currMax
	}
	fmt.Println(res)
}
