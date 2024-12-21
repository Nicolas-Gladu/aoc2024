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

	part1 := solutionPart1(challenge.Inputs)
	assert.Equal(t, 18, part1)
	part2 := solutionPart2(challenge.Inputs)
	assert.Equal(t, 9, part2)
}
