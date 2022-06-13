package tree

import (
	"demo/util"
	"testing"
)

func Test_mergeTrees(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			args: args{
				root1: getTreeNode(),
				root2: getMirrorTreeNode(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTrees(tt.args.root1, tt.args.root2)
			t.Logf("mergeTrees() = %v", util.MarshalIndent(got))
		})
	}
}
