package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sssdrum/aoc-2025/utils"
)

func main() {
	part1()
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
