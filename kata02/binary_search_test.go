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
	}, {
		num:      59,
		arr:      []int{-53, -59, 0, 53},
		expected: -1,
	}, {
		num:      39,
		arr:      []int{},
		expected: -1,
	}, {
		num:      493,
		arr:      []int{493},
		expected: 0,
	}, {
		num:      -48,
		arr:      []int{-48},
		expected: 0,
	}, {
		num:      0,
		arr:      []int{0, 3},
		expected: 0,
	}, {
		num:      5,
		arr:      []int{-3, 0, 5},
		expected: 2,
	}, {
		num:      18,
		arr:      []int{3, 0, 9},
		expected: -1,
	}, {
		num:      11,
		arr:      []int{11, 33, 100},
		expected: 0,
	}, {
		num:      43,
		arr:      []int{43, 66, 122, 655},
		expected: 0,
	}, {
		num:      -52,
		arr:      []int{43, 66, 122, 675},
		expected: -1,
	}, {
		num:      -2,
		arr:      []int{-66, -43, -12, -2},
		expected: 3,
	}, {
		num:      -2,
		arr:      []int{-143, -66, -2, 52},
		expected: 2,
	},
}

func TestIterativeSearch(t *testing.T) {
	require := require.New(t)

	for _, test := range tests {
		v := kata02.IterativeSearch(test.num, test.arr)
		require.Equal(test.expected, v)
	}
}

func TestRecursiveSearch(t *testing.T) {
	require := require.New(t)

	for _, test := range tests {
		v := kata02.RecursiveSearch(test.num, test.arr)
		require.Equal(test.expected, v)
	}
}
