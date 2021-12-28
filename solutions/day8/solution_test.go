package day8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIt(t *testing.T) {
	output := process("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf")
	assert.NotNil(t, output)
}
