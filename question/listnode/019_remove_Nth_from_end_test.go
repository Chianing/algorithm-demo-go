package listnode

import (
	"demo/util"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				head: getCustomListNode(1, 2),
				n:    1,
			},
		},
		{
			args: args{
				head: getCustomListNode(1, 2, 3),
				n:    3,
			},
		},
		{
			args: args{
				head: getCustomListNode(1, 2, 3),
				n:    2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("removeNthFromEnd() done, result: %v", util.MarshalIndent(removeNthFromEnd(tt.args.head, tt.args.n)))
		})
	}
}
