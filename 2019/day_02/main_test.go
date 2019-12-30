package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_exec(t *testing.T) {
	tests := []struct {
		name string
		ops  []int
		want []int
	}{
		{ops: []int{1, 0, 0, 0, 99}, want: []int{2, 0, 0, 0, 99}},
		{ops: []int{2, 3, 0, 3, 99}, want: []int{2, 3, 0, 6, 99}},
		{ops: []int{2, 4, 4, 5, 99, 0}, want: []int{2, 4, 4, 5, 99, 9801}},
		{ops: []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, want: []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, exec(tt.ops))
		})
	}
}

func Test_part1(t *testing.T) {
	tests := []struct {
		name             string
		ops              []int
		noun, verb, want int
	}{
		{ops: input, noun: 12, verb: 2, want: 3716250},
		{ops: input, noun: 64, verb: 72, want: 19690720}, // answer: 100*noun+verb=6472
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ops := make([]int, len(tt.ops))
			copy(ops, tt.ops)
			ops[1] = tt.noun
			ops[2] = tt.verb

			assert.NotEqual(t, tt.ops, ops)

			assert.Equal(t, tt.want, exec(ops)[0])
		})
	}
}
