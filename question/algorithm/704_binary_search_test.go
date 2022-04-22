package algorithm

import "testing"

func Test_binarySearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	nums := []int{1, 2, 3, 4}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums:   nums,
				target: 1,
			},
			want: 0,
		},
		{
			args: args{
				nums:   nums,
				target: 4,
			},
			want: 3,
		},
		{
			args: args{
				nums:   nums,
				target: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
