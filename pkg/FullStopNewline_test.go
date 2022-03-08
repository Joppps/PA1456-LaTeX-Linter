package Lint

import "testing"

func TestNewLiner(t *testing.T) {
	type args struct {
		line        string
		nrOfIndents int
		preString   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"first", args{"dankmems. Testing. %Comment", 0, ""}, "dankmems. \nTesting. %Comment"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLiner(tt.args.line, tt.args.nrOfIndents, tt.args.preString); got != tt.want {
				t.Errorf("NewLiner() = %v, want %v", got, tt.want)
			}
		})
	}
}
