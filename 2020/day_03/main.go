package main

import "fmt"

const (
	square = '.'
	tree   = '#'
)

func main() {
	answer1 := findPath(3, 1, len(input[0]))
	fmt.Println("part 1:", answer1) // 151

	answer2 := findPath(1, 1, len(input[0]))
	answer2 *= findPath(3, 1, len(input[0]))
	answer2 *= findPath(5, 1, len(input[0]))
	answer2 *= findPath(7, 1, len(input[0]))
	answer2 *= findPath(1, 2, len(input[0]))

	fmt.Println("part 2:", answer2) // 7540141059
}

func findPath(xi, yi, width int) int {
	var count int

	for x, y := 0, 0; y < len(input); {
		if input[y][x] == tree {
			count++
		}

		x = (x + xi) % width
		y += yi
	}

	return count
}
