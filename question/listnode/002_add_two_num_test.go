package listnode

import (
	"demo/util"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "boundary case",
			args: args{
				l1: nil,
				l2: nil,
			},
		},
		{
			name: "case1",
			args: args{
				l1: getSingleListNode(),
				l2: getListNode(),
			},
		},
		{
			name: "case2",
			args: args{
				l1: getListNode(),
				l2: getListNode(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("after addTwoNumbers(), args: %v", util.MarshalIndent(addTwoNumbers(tt.args.l1, tt.args.l2)))
		})
	}
}
