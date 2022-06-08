package tree

import "testing"

func Test_longestZigZag(t *testing.T) {
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
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestZigZag(tt.args.root); got != tt.want {
				t.Errorf("longestZigZag() = %v, want %v", got, tt.want)
			}
		})
	}
}
