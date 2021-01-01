package main

import "fmt"

func main() {
	result1 := part1()
	fmt.Println("answer 1:", result1) // 467

	result2 := part2()
	fmt.Println("answer 2:", result2) // 441
}

func part1() int {
	var result int

	for i := range input {
		var count int

		for r := range input[i].pswrd {
			if input[i].pswrd[r] == byte(input[i].word) {
				count++
			}
		}

		if count <= input[i].max && count >= input[i].min {
			result++
		}
	}

	return result
}

func part2() int {
	var result int

	for i := range input {
		if input[i].pswrd[input[i].min-1] == byte(input[i].word) && input[i].pswrd[input[i].max-1] != byte(input[i].word) {
			result++
		} else if input[i].pswrd[input[i].min-1] != byte(input[i].word) && input[i].pswrd[input[i].max-1] == byte(input[i].word) {
			result++
		}
	}

	return result
}
