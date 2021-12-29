package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		line   string
		result result
		score  int
	}{{
		line:   "<>",
		result: resultOk,
		score:  0,
	}, {
		line:   "{([(<{}[<>[]}>{[]{[(<()>",
		result: resultCorrupted,
		score:  1197,
	}}

	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {
			score, result := parse(tt.line)
			assert.Equal(t, tt.score, score)
			assert.Equal(t, tt.result, result)
		})
	}
}
