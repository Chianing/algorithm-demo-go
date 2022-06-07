package listnode

import (
	"demo/util"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				head:  getCustomListNode(1, 2, 3, 4, 5),
				left:  2,
				right: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseBetween(tt.args.head, tt.args.left, tt.args.right)
			t.Logf("reverseBetween() = %v", util.MarshalIndent(got))
		})
	}
}

func Test_reverseN(t *testing.T) {
	type args struct {
		root *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				root: getCustomListNode(1, 2, 3, 4),
				n:    3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseN(tt.args.root, tt.args.n)
			t.Logf("reverseN() = %v", util.MarshalIndent(got))
		})
	}
}

func Test_reverseBetween2(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				head:  getCustomListNode(1, 2, 3, 4, 5),
				left:  2,
				right: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseBetween2(tt.args.head, tt.args.left, tt.args.right)
			t.Logf("reverseBetween2() = %v", util.MarshalIndent(got))
		})
	}
}
