package main

import (
	"Majha/AdvantOfCode/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	challenge, err := pkg.NewDayChallenge("test.txt")
	if err != nil {
		panic(err)
	}

	part1, part2 := solution(challenge.Inputs)
	assert.Equal(t, 2, part1)
	assert.Equal(t, 4, part2)
}

func BenchmarkSolution(b *testing.B) {
	// Generate test data
	challenge, err := pkg.NewDayChallenge("input.txt")
	if err != nil {
		panic(err)
	}

	b.ResetTimer() // Reset the timer to exclude setup time
	for i := 0; i < b.N; i++ {
		_, _ = solution(challenge.Inputs)
	}
}
