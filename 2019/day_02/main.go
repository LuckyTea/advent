package main

import "fmt"

const (
	sum = 1
	mul = 2
	fin = 99
)

func main() {
	fmt.Println("https://adventofcode.com/2019/day/2")
}

func exec(ops []int) []int {
	for i := 0; i < len(ops); i++ {
		switch ops[i] {
		case sum:
			ops[ops[i+3]] = ops[ops[i+1]] + ops[ops[i+2]]
			i += 3
		case mul:
			ops[ops[i+3]] = ops[ops[i+1]] * ops[ops[i+2]]
			i += 3
		case fin:
			return ops
		}
	}

	return ops
}
