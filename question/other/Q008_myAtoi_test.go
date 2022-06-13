package other

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	args: args{s: "-123"},
		// 	want: -123,
		// },
		// {
		// 	args: args{s: "123s"},
		// 	want: 123,
		// },
		{
			args: args{s: " -0012a42"},
			want: -12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
