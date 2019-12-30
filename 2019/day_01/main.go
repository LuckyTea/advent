package main

import "fmt"

func main() {
	fmt.Println("https://adventofcode.com/2019/day/1")
}

func count(mass int) int {
	return mass/3 - 2
}

func part1(input []int) (total int) {
	for i := range input {
		total += count(input[i])
	}

	return
}

func part2(input []int) (total int) {
	for i := range input {
		tmp := count(input[i])

		for {
			if tmp > 0 {
				total += tmp
				tmp = count(tmp)

				continue
			}

			break
		}
	}

	return
}
