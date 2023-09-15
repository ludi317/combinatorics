package itertools

import (
	"reflect"
	"testing"
)

func TestCartesianPower(t *testing.T) {
	type args struct {
		size uint64
		n    int
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			name: "CartesianPower(2, 3)",
			args: args{
				size: 2,
				n:    3,
			},
			want: [][]byte{
				{0, 0, 0},
				{0, 0, 1},
				{0, 1, 0},
				{0, 1, 1},
				{1, 0, 0},
				{1, 0, 1},
				{1, 1, 0},
				{1, 1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CartesianPower(tt.args.size, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartesianPower() = %v, want %v", got, tt.want)
			}
		})
	}
}
