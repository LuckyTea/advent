package main

import (
	"fmt"
	"sort"
)

func main() {
	answer1 := part1()
	fmt.Println("part 1 answer:", answer1)

	answer2 := part2()
	fmt.Println("part 2 answer:", answer2)
}

const want = 2020

func part1() int {
	sort.Ints(input)

	var (
		got int
		b   = 0
		e   = len(input) - 1
	)

	for {
		got = input[b] + input[e]

		switch {
		case got == want:
		case got > want:
			e -= 1
			continue
		case got < want:
			b += 1
			continue
		}

		break
	}

	return input[b] * input[e] // 355875
}

func part2() int {
	sort.Ints(input)

	for i := range input {
		for ii := range input {
			for iii := range input {
				if input[i]+input[ii]+input[iii] == want {
					return input[i] * input[ii] * input[iii]
				}
			}
		}
	}

	return 0
}
