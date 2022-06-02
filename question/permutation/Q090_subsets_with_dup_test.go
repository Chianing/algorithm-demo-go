package permutation

import (
	"demo/util"
	"testing"
)

func Test_subsetsWithDup(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				nums: []int{1, 2, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subsetsWithDup(tt.args.nums)
			t.Logf("subsetsWithDup() = %v", util.MarshalIndent(got))
		})
	}
}
