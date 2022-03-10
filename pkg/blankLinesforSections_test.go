package Lint

import "testing"

func TestInsertBlankLinesOnSections(t *testing.T) {
	type args struct {
		line      string
		lineCount int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"first", args{"Nothing here", 10}, "Nothing here"},
		{"second", args{"\\chapter", 2}, "\n\n\\chapter"},
		{"third", args{"\\section{", 3}, "\n\n\n\\section{"},
		{"fourth", args{"\\subsubsection", 2}, "\n\n\\subsubsection"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertBlankLinesOnSections(tt.args.line, tt.args.lineCount); got != tt.want {
				t.Errorf("InsertBlankLinesOnSections() = %v, want %v", got, tt.want)
			}
		})
	}
}
