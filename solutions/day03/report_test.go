package day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReport(t *testing.T) {
	tests := []struct {
		name     string
		fn       func(r Report) (int64, error)
		expected int
	}{{
		name: "gamma",
		fn: func(r Report) (int64, error) {
			return r.Gamma()
		},
		expected: 22,
	}, {
		name: "epsilon",
		fn: func(r Report) (int64, error) {
			return r.Epsilon()
		},
		expected: 9,
	}, {
		name: "oxygen",
		fn: func(r Report) (int64, error) {
			return r.Oxygen()
		},
		expected: 23,
	}, {
		name: "co2",
		fn: func(r Report) (int64, error) {
			return r.C02()
		},
		expected: 10,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			report, err := NewReport("testdata.txt")
			assert.NoError(t, err)

			actual, err := tt.fn(report)
			assert.NoError(t, err)

			assert.Equal(t, tt.expected, int(actual))
		})
	}
}
