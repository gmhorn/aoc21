package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	data := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	assert.Equal(t, 7, reduce(data, increase))
}

func TestWindowSum(t *testing.T) {
	data := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	expected := []int{607, 618, 618, 617, 647, 716, 769, 792}
	assert.Equal(t, expected, windowSum(data, 3))
}
