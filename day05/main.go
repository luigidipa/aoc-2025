package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/luigidipa/aoc-2025/utils"
)

func main() {
	part1()
	part2()
}

func part1() {
	lines := utils.ReadLines("./input.txt")
	// Find ranges of fresh ingredient ranges and available ingredients
	res := 0
	var ranges []string
	var ingrs []string
	for i, l := range lines {
		if l == "" {
			ranges = lines[0:i]
			ingrs = lines[i+1:]
			break
		}
	}

	for _, i := range ingrs {
		id, _ := strconv.Atoi(i)
		for _, r := range ranges {
			toks := strings.Split(r, "-")
			start, _ := strconv.Atoi(toks[0])
			end, _ := strconv.Atoi(toks[1])
			if start <= id && id <= end {
				res++
				break
			}
		}
	}
	fmt.Println(res)
}

func part2() {
	res := 0
	lines := utils.ReadLines("./input.txt")
	size := 0
	for lines[size] != "" {
		size++
	}

	ranges := make([][]int, size)
	for i := range size {
		toks := strings.Split(lines[i], "-")
		start, _ := strconv.Atoi(toks[0])
		end, _ := strconv.Atoi(toks[1])
		ranges[i] = []int{start, end}
	}
	// Sort ranges on start time
	sort.Slice(ranges, func(a, b int) bool {
		return ranges[a][0] < ranges[b][0]
	})

	ps, pe := -1, -1
	for _, r := range ranges {
		s, e := r[0], r[1]
		if s > pe { // no overlap
			res += e - s + 1
		} else if s == pe { // range starts where prev range ends
			res += e - s
		} else if s >= ps && e > pe { // overlap
			res += e - pe
		}
		ps = s
		if e > pe { // Only update end if greater than previous end. Not checking will break the algo
			pe = e
		}
	}

	fmt.Println(res)
}
