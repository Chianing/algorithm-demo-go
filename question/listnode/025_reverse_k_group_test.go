package listnode

import (
	"demo/util"
	"testing"
)

func Test_reverseByK1(t *testing.T) {
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
				getCustomListNode(1, 2, 3, 4),
				2,
			},
		},
		{
			args: args{
				getCustomListNode(1, 2, 3, 4),
				3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("reverseByK1() done, got: %v", util.MarshalIndent(reverseByK1(tt.args.head, tt.args.k)))
		})
	}
}

func Test_reverseKGroup(t *testing.T) {
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
				head: getCustomListNode(1, 2, 3, 4),
				k:    2,
			},
		},
		{
			args: args{
				head: getCustomListNode(1, 2, 3),
				k:    2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("reverseByK2() done, got: %v", util.MarshalIndent(reverseKGroup(tt.args.head, tt.args.k)))
		})
	}
}
