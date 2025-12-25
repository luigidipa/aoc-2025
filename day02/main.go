package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/sssdrum/aoc-2025/utils"
)

type idRange struct {
	start int
	end   int
}

func main() {
	part1()
	part2()
}

func part1() {
	res := 0
	input, _ := utils.ReadInput("./input.txt")
	ranges := getRanges(input)
	for _, r := range ranges {
		for i := r.start; i <= r.end; i++ {
			s := strconv.Itoa(i)
			length := len(s)
			if length%2 != 0 {
				continue
			}
			left := s[0 : length/2]
			right := s[length/2:]
			if left == right {
				res += i
			}
		}
	}
	fmt.Println(res)
}

func part2() {
	res := 0
	input, _ := utils.ReadInput("./input.txt")
	ranges := getRanges(input)
	for _, r := range ranges {
		for i := r.start; i <= r.end; i++ {
			if isInvalidId(i) {
				res += i
			}
		}
	}
	fmt.Println(res)
}

func getRanges(input string) []idRange {
	toks := strings.Split(strings.Trim(input, "\n"), ",")
	ranges := make([]idRange, len(toks))
	for i, t := range toks {
		parts := strings.Split(t, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		ranges[i] = idRange{
			start: start,
			end:   end,
		}
	}
	return ranges
}

func isInvalidId(n int) bool {
	s := strconv.Itoa(n)
	length := len(s)
	for i := range length / 2 {
		pattern := s[0 : i+1]
		regex := "^(" + regexp.QuoteMeta(pattern) + ")+$"
		match, _ := regexp.MatchString(regex, s)
		if match {
			return true
		}
	}
	return false
}
