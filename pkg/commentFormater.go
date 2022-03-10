package Lint

import (
	"regexp"
)

func CommentFormat(line string, tabsOrNot bool, nrOfSpacers int) string {
	var spacer string
	var newLine string = line
	for i := 0; i < nrOfSpacers; i++ {
		if tabsOrNot {
			spacer += "\t"
		} else {
			spacer += " "
		}
	}

	re := regexp.MustCompile(`(?m)([^\\](%))`) //find all "%" except "\%" and replace them with the "%spacer"
	newLine = re.ReplaceAllString(newLine, "$1"+spacer)
	return newLine
}
