package permutation

import (
	"demo/util"
	"testing"
)

func Test_combinationSum3(t *testing.T) {
	type args struct {
		k int
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				k: 3,
				n: 7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum3(tt.args.k, tt.args.n)
			t.Logf("combinationSum3() = %v", util.MarshalIndent(got))
		})
	}
}
