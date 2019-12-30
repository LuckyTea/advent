package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_count(t *testing.T) {
	tests := []struct {
		name string
		mass int
		want int
	}{
		{mass: 12, want: 2},
		{mass: 14, want: 2},
		{mass: 1969, want: 654},
		{mass: 100756, want: 33583},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, count(tt.mass))
		})
	}
}

func Test_part1(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		wantTotal int
	}{
		{input: input, wantTotal: 3342050},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantTotal, part1(tt.input))
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		wantTotal int
	}{
		{input: []int{14}, wantTotal: 2},
		{input: []int{1969}, wantTotal: 966},
		{input: []int{100756}, wantTotal: 50346},
		{input: input, wantTotal: 5010211},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				assert.Equal(t, tt.wantTotal, part2(tt.input))
			})
		})
	}
}
