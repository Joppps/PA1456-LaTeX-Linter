package Lint

import "testing"

func TestCommentFormat(t *testing.T) {
	type args struct {
		line        string
		tabsOrNot   bool
		nrOfSpacers int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"first", args{"test%Comment%Test", true, 1}, "test%\tComment%\tTest"},
		{"Second", args{"math \\%percentage %This is good! %comment", false, 3}, "math \\%percentage %   This is good! %   comment"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CommentFormat(tt.args.line, tt.args.tabsOrNot, tt.args.nrOfSpacers); got != tt.want {
				t.Errorf("CommentFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
