package algorithm

import (
	"reflect"
	"testing"
)

var arr1, arr1Want []int
var arr2, arr2Want []int

func init() {
	arr1 = []int{5, 2, 3, 1}
	arr1Want = []int{1, 2, 3, 5}

	arr2 = []int{8, 9, 1, 7, 2, 3, 5, 4, 6, 0}
	arr2Want = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
}

func Test_sortBubble(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{arr: arr1},
			want: arr1Want,
		},
		{
			args: args{arr: arr2},
			want: arr2Want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortBubble(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortBubble() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortInsert(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{arr: arr1},
			want: arr1Want,
		},
		{
			args: args{arr: arr2},
			want: arr2Want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortInsert(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortShell(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{arr: arr1},
			want: arr1Want,
		},
		{
			args: args{arr: arr2},
			want: arr2Want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortShell(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortShell() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortMerge(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{arr: arr1},
			want: arr1Want,
		},
		{
			args: args{arr: arr2},
			want: arr2Want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortMerge(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortQuick(t *testing.T) {
	type args struct {
		arr   []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				arr:   arr1,
				left:  0,
				right: len(arr1) - 1,
			},
			want: arr1Want,
		},
		{
			args: args{
				arr:   arr2,
				left:  0,
				right: len(arr2) - 1,
			},
			want: arr2Want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortQuick(tt.args.arr, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortQuick() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_heapSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{arr: arr1},
			want: arr1Want,
		},
		{
			args: args{arr: arr2},
			want: arr2Want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := heapSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("heapSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
