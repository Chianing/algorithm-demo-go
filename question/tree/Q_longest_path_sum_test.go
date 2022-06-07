package tree

import "testing"

func Test_longestPathSum1(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				root: getTreeNode(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestPathSum1(tt.args.root)
			t.Logf("longestPathSum1() = %v", got)
		})
	}
}
