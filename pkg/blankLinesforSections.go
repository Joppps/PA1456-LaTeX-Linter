package Lint

import (
	"regexp"
)

//inserting the a number of blanklines before the section/chapters...
func InsertBlankLinesOnSections(line string, lineCount int) string {
	re := regexp.MustCompile(`\r`)
	var windows bool = re.MatchString(line) //returns true if carriage return exists.

	var blanks string
	for i := 0; i < lineCount; i++ {
		if windows {
			blanks += "\r\n"
		} else {
			blanks += "\n"
		}
	}
	re = regexp.MustCompile(`(?m)(^\s*(\\section|\\subsection|\\subsubsection|\\chapter|\\paragraph|\\subparagraph))`)
	var newLine string = re.ReplaceAllString(line, blanks+"$1")
	return newLine
}
