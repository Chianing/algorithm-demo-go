package listnode

import (
	"demo/util"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				getCustomListNode(0, 0, 0, 0, 0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("deleteDuplicates() done, got: %v", util.MarshalIndent(deleteDuplicates(tt.args.head)))
		})
	}
}
