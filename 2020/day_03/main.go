package main

import "fmt"

const (
	square = '.'
	tree   = '#'
)

func main() {
	answer1 := part1()
	fmt.Println("part 1:", answer1) // 151
}

func part1() int {
	var (
		count int
		width = len(input[0])
	)

	for x, y := 0, 0; y < len(input); {
		if input[y][x] == tree {
			count++
		}

		x = (x + 3) % width
		y += 1
	}

	return count
}
