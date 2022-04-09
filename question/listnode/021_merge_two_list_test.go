package listnode

import (
	"demo/util"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// {
		// 	args: args{
		// 		list1: getCustomListNode(1, 3, 5),
		// 		list2: getCustomListNode(2, 3, 6),
		// 	},
		// },
		{
			args: args{
				list1: getCustomListNode(1, 2, 4),
				list2: getCustomListNode(1, 3, 4),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("mergeTwoList() done, ret: %v", util.MarshalIndent(mergeTwoLists(tt.args.list1, tt.args.list2)))
		})
	}
}
