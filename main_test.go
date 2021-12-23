package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWindowSum(t *testing.T) {
	data := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	expected := []int{607, 618, 618, 617, 647, 716, 769, 792}
	actual := WindowSum(data, 3)
	assert.Equal(t, expected, actual)
}

func TestReduce(t *testing.T) {
	data := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	expected := 7
	actual := Reduce(data, Increase)
	assert.Equal(t, expected, actual)
}
