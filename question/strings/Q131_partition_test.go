package strings

import (
	"demo/util"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// {
		// 	args: args{s: "aab"},
		// },
		// {
		// 	args: args{s: "a"},
		// },
		{
			args: args{
				s: "cbbbcc",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := partition(tt.args.s)
			t.Logf("partition() = %v", util.MarshalIndent(got))
		})
	}
}
