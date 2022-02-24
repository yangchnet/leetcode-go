package offer04

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findNumberIn2DArray(t *testing.T) {
	require.Equal(t, true, findNumberIn2DArray([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5))

	require.Equal(t, false, findNumberIn2DArray([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20))

	require.Equal(t, false, findNumberIn2DArray([][]int{{}, {}, {}, {}, {}}, -1))
}
