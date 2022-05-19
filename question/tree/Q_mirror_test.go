package tree

import "testing"

func Test_isMirror(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				root1: getTreeNode(),
				root2: getTreeNode(),
			},
			want: false,
		},
		{
			args: args{
				root1: getTreeNode(),
				root2: getMirrorTreeNode(),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMirror(tt.args.root1, tt.args.root2); got != tt.want {
				t.Errorf("isMirror() = %v, want %v", got, tt.want)
			}
		})
	}
}
