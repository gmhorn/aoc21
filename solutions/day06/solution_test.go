package day06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestState(t *testing.T) {
	var init state
	init[1] = 1
	init[2] = 1
	init[3] = 2
	init[4] = 1

	day1 := nextState(init)
	assert.Equal(t, state([9]int{1, 1, 2, 1, 0, 0, 0, 0, 0}), day1)
	day2 := nextState(day1)
	assert.Equal(t, state([9]int{1, 2, 1, 0, 0, 0, 1, 0, 1}), day2)
	day3 := nextState(day2)
	assert.Equal(t, state([9]int{2, 1, 0, 0, 0, 1, 1, 1, 1}), day3)

	next := init
	for i := 0; i < 18; i++ {
		next = nextState(next)
	}
	assert.Equal(t, 26, sum(next))
}
