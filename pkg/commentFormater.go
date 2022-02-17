package Lint

import "strings"

func CommentFormat(line string, tabsOrNot bool, nrOfSpaces int) string {
	var spacer string
	var newLine string = line
	for i := 0; i < nrOfSpaces; i++ {
		if tabsOrNot {
			spacer += "\t"
		} else {
			spacer += " "
		}
	}
	if strings.Contains(line, "%") {
		if !strings.Contains(line, "% ") { //dont want to format correct formatting.
			newLine = strings.Replace(line, "%", "%"+spacer, -1)
		}
	}
	return newLine
}
