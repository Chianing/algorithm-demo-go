package test

import "testing"

func Test_routinePoolExecute(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			routinePoolExecute()
		})
	}
}

func Test_printInOrder(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printInOrder()
		})
	}
}
