package main

import (
	"fmt"
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
	ops := strings.Fields(lines[len(lines)-1])
	tots := make([]int, len(ops))
	for i, o := range ops {
		if o == "*" {
			tots[i] = 1
		}
	}

	for _, l := range lines[:len(lines)-1] {
		fields := strings.Fields(l)
		for i, f := range fields {
			n, _ := strconv.Atoi(f)
			if ops[i] == "*" {
				tots[i] *= n
			} else {
				tots[i] += n
			}
		}
	}

	res := 0
	for _, t := range tots {
		res += t
	}
	fmt.Println(res)
}

func part2() {
	lines := utils.ReadLines("./input.txt")

	// Add whitespace to shorter lines
	max_len := 0
	for _, l := range lines {
		max_len = max(max_len, len(l))
	}
	for i, l := range lines {
		if len(l) < max_len {
			lines[i] = lines[i] + strings.Repeat(" ", max_len-len(l))
		}
	}

	// Kinda horrible but it does the job
	num_lines := len(lines)
	last_line := string(lines[num_lines-1])
	res := 0
	for i := range len(last_line) {
		if last_line[i] != ' ' {
			op := last_line[i]
			tot := 0
			if op == '*' {
				tot = 1
			}
			j := i
			for j < max_len-1 && last_line[j+1] == ' ' || j == max_len-1 {
				s := ""
				for _, l := range lines[0 : num_lines-1] {
					s += string(l[j])
				}
				j++
				n, _ := strconv.Atoi(strings.TrimSpace(s))
				if op == '*' {
					tot *= n
				} else {
					tot += n
				}
			}
			res += tot
		}
	}
	fmt.Println(res)
}
