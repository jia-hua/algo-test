package world

import (
	"testing"
)

func TestNextCellStep(t *testing.T) {
	type args struct {
		subGrid []bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"given dead, 3 above living, then should born", args{subGrid: []bool{
			true, true, true, false, false, false, false, false, false,
		}}, true},
		{"given all dead, no living, should stay dead", args{subGrid: []bool{
			false, false, false, false, false, false, false, false, false,
		}}, false},
		{"given living, 3 above living, then should stay", args{subGrid: []bool{
			true, true, true, false, true, false, false, false, false,
		}}, true},
		{"given living, 3 neighbor living, then should stay", args{subGrid: []bool{
			true, true, false, false, true, false, false, false, true,
		}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NextCellStep(tt.args.subGrid); got != tt.want {
				t.Errorf("NextCellStep() = %v, want %v", got, tt.want)
			}
		})
	}
}
