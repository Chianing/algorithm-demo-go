package tree

import "testing"

func Test_findTilt(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{root: getTreeNode()},
			want: 41,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTilt(tt.args.root); got != tt.want {
				t.Errorf("findTilt() = %v, want %v", got, tt.want)
			}
		})
	}
}
