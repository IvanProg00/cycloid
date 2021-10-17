package kata02_test

import (
	"testing"

	"github.com/IvanProg00/cycloid/kata02"
	"github.com/stretchr/testify/require"
)

var tests = []struct {
	num      int
	arr      []int
	expected int
}{
	{
		num:      1,
		arr:      []int{0, 1, 2},
		expected: 1,
	},
	{
		num:      59,
		arr:      []int{-53, 0, -59, 53},
		expected: -1,
	},
}

func TestIterativeSearch(t *testing.T) {
	for _, test := range tests {
		v := kata02.IterativeSearch(test.num, test.arr)
		require.Equal(t, v, test.expected)
	}
}
