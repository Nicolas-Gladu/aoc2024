package main

import (
	"Majha/AdvantOfCode/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	challenge, err := pkg.NewDayChallenge("test.txt")
	if err != nil {
		panic(err)
	}
	part1, part2 := solution(challenge)
	assert.Equal(t, 11, part1)
	assert.Equal(t, 31, part2)
}
