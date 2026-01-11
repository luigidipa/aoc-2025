package main

import (
	"fmt"
	"strconv"

	"github.com/luigidipa/aoc-2025/utils"
)

func main() {
	part1()
	part2()
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

func part2() {
	lines := utils.ReadLines("./input.txt")
	res := 0
	lineLen := len(lines[0])
	for _, line := range lines {
		digits, need, start, end, s := lineLen, 12, 0, 0, ""
		for need > 0 {
			end = start + digits - need
			maxFound, pos := maxIntInStr(line[start : end+1])
			s += strconv.Itoa(maxFound)
			start += pos + 1
			need--
			digits = len(line) - start
		}
		n, _ := strconv.Atoi(s)
		res += n
	}
	fmt.Println(res)
}

// Assumes the input is composed of only positive integers. Returns the maximum
// integer found, with its position in the string
func maxIntInStr(s string) (int, int) {
	maxFound, pos := -1, 0
	for i, r := range s {
		n, _ := strconv.Atoi(string(r))
		if n > maxFound {
			maxFound = n
			pos = i
		}
	}
	return maxFound, pos
}
