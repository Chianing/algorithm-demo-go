package other

import "testing"

func Test_getLongestCost(t *testing.T) {
	type args struct {
		nodes []*invokeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nodes: []*invokeNode{
					{
						startName: "A",
						endName:   "B",
						cost:      100,
					},
					{
						startName: "A",
						endName:   "C",
						cost:      200,
					},
					{
						startName: "B",
						endName:   "C",
						cost:      150,
					},
					{
						startName: "D",
						endName:   "E",
						cost:      300,
					},
				},
			},
			want: 300,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLongestCost(tt.args.nodes); got != tt.want {
				t.Errorf("getLongestCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
