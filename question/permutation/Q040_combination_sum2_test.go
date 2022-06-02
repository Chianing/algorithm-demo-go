package permutation

import (
	"demo/util"
	"testing"
)

func Test_combinationSum2(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				candidates: []int{2, 5, 2, 1, 2},
				target:     5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum2(tt.args.candidates, tt.args.target)
			t.Logf("combinationSum2() = %v", util.MarshalIndent(got))
		})
	}
}
