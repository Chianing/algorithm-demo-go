package listnode

import (
	"demo/util"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{head: getCustomListNode(1, 2, 3, 4)},
			want: getCustomListNode(2, 1, 4, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := swapPairs(tt.args.head)
			t.Logf("swapPairs() = %v", util.MarshalIndent(got))
		})
	}
}
