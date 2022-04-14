package listnode

import (
	"demo/util"
	"testing"
)

func Test_rotateRight(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				head: getCustomListNode(1, 2, 3, 4, 5),
				k:    2,
			},
		},
		{
			args: args{
				head: getCustomListNode(0, 1, 2),
				k:    4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("rotateRight() done, got: %v", util.MarshalIndent(rotateRight(tt.args.head, tt.args.k)))
		})
	}
}
