package other

import (
	"testing"
)

func Test_threeSum(t *testing.T) {
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
				nums: []int{-1, 0, 1, 2, -1, -4},
			},
		},
		{
			args: args{
				nums: []int{-2, 0, 0, 2, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.args.nums)
			t.Logf("threeSum() = %v", got)
		})
	}
}
