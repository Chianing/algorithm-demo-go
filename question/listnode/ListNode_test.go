package listnode

import (
	"demo/util"
	"testing"
)

func Test_getCustomListNode(t *testing.T) {
	type args struct {
		vals []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				[]int{1, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("before getCustomListNode(), args: %v", tt.args)
			t.Logf("after getCustomListNode(), args: %v", util.MarshalIndent(getCustomListNode(tt.args.vals...)))
		})
	}
}
