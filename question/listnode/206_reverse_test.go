package listnode

import (
	"demo/util"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "boundary case",
			args: args{
				nil,
			},
		},
		{
			name: "single case",
			args: args{
				getSingleListNode(),
			},
		},
		{
			name: "normal case",
			args: args{
				getListNode(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("before reverseList() = %v", util.MarshalIndent(tt.args.head))
			t.Logf("finish reverseList() = %v", util.MarshalIndent(reverseList1(tt.args.head)))
		})
	}
}

func Test_reverseList2(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "boundary case",
			args: args{
				nil,
			},
		},
		{
			name: "single case",
			args: args{
				getSingleListNode(),
			},
		},
		{
			name: "normal case",
			args: args{
				getListNode(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("before reverseList() = %v", util.MarshalIndent(tt.args.head))
			t.Logf("finish reverseList() = %v", util.MarshalIndent(reverseList2(tt.args.head)))
		})
	}
}
