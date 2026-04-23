package medium

import (
	"dsa/structs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_cloneGraph(t *testing.T) {
	t.Parallel()
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			args: args{node: structs.BuildGraph([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}})},
			want: structs.BuildGraph([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, structs.PrintGraph(tt.want), structs.PrintGraph(cloneGraph(tt.args.node)), "graph structure mismatch")
		})
	}
}
