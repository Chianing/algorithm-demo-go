package tree

import (
	"reflect"
	"testing"
)

func Test_zigzagLevelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantRet [][]int
	}{
		{
			args: args{
				getTreeNode(),
			},
			wantRet: [][]int{
				{3},
				{20, 9},
				{15, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := zigzagLevelOrder(tt.args.root); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("zigzagLevelOrder() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
