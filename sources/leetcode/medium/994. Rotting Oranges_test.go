package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_orangesRotting(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			// 2 1 1
			// 1 1 0
			// 0 1 1
			name: "standard case all reachable",
			grid: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			want: 4,
		},
		{
			// 2 1 1
			// 0 1 1
			// 1 0 1  <- bottom-left fresh is isolated
			name: "unreachable fresh orange returns -1",
			grid: [][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			want: -1,
		},
		{
			// 0 2
			// no fresh oranges at all
			name: "no fresh oranges returns 0",
			grid: [][]int{
				{0, 2},
			},
			want: 0,
		},
		{
			// 1 1
			// 1 1  <- no rotten, all fresh
			name: "no rotten oranges fresh present returns -1",
			grid: [][]int{
				{1, 1},
				{1, 1},
			},
			want: -1,
		},
		{
			// 1 0 1  <- fresh separated by empty cell from rotten
			// 2 0 0
			name: "fresh separated by empty cell returns -1",
			grid: [][]int{
				{1, 0, 1},
				{2, 0, 0},
			},
			want: -1,
		},
		{
			// 2 2
			// 2 2  <- all rotten from start
			name: "all rotten from start returns 0",
			grid: [][]int{
				{2, 2},
				{2, 2},
			},
			want: 0,
		},
		{
			// 2 1 1
			// 1 1 1
			// 1 1 1  <- rotten at top-left corner spreads through whole grid
			name: "single rotten in corner spreads through full grid",
			grid: [][]int{
				{2, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			want: 4,
		},
		{
			// 1x1 empty cell
			name: "1x1 empty cell returns 0",
			grid: [][]int{{0}},
			want: 0,
		},
		{
			// 1x1 fresh orange, no rotten
			name: "1x1 fresh orange returns -1",
			grid: [][]int{{1}},
			want: -1,
		},
		{
			// 1x1 rotten orange
			name: "1x1 rotten orange returns 0",
			grid: [][]int{{2}},
			want: 0,
		},
		{
			// 2 1 1 1 1  <- single row, rotten at left edge
			name: "single row rotten at left propagates right",
			grid: [][]int{
				{2, 1, 1, 1, 1},
			},
			want: 4,
		},
		{
			// two rotten sources speed up rotting
			// 2 1 1 1 2
			name: "two rotten sources both ends meet in middle",
			grid: [][]int{
				{2, 1, 1, 1, 2},
			},
			want: 2,
		},
		{
			// 0 0 0
			// 0 0 0  <- fully empty grid
			name: "fully empty grid returns 0",
			grid: [][]int{
				{0, 0, 0},
				{0, 0, 0},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := orangesRotting(tt.grid)
			assert.Equal(t, tt.want, got)
		})
	}
}
