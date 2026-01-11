package main

import (
	"fmt"
	"strconv"

	"github.com/luigidipa/aoc-2025/utils"
)

const LEN = 100

func main() {
	part1()
	part2()
}

func part1() {
	dial, res := 50, 0
	lines := utils.ReadLines("./input.txt")
	for _, l := range lines {
		dir, tok := l[0], l[1:]
		steps, err := strconv.Atoi(tok)
		utils.CheckErr(err)

		if dir == 'R' {
			dial = (dial + steps) % (LEN)
		} else {
			dial = ((dial - steps) + LEN) % (LEN)
		}

		if dial == 0 {
			res++
		}
	}
	fmt.Println(res)
}

func part2() {
	dial, res := 50, 0
	lines := utils.ReadLines("./input.txt")
	for _, l := range lines {
		dir, tok := l[0], l[1:]
		steps, err := strconv.Atoi(tok)
		utils.CheckErr(err)

		if dir == 'R' {
			res += (dial + steps) / LEN
			dial = (dial + steps) % LEN
		} else {
			if dial == 0 {
				res += steps / LEN
			} else if steps >= dial {
				res += 1 + (steps-dial)/LEN
			}
			dial = ((dial - steps%LEN) + LEN) % LEN
		}
	}
	fmt.Println(res)
}
