package Lint

import (
	"regexp"
	"strings"
)

func NewLiner(line string, nrOfIndents int, preString string) string {
	//split the string into two and then make the fullstop thingy on the non comment part.
	//then add the comment back onto the newLine added string.

	lines := strings.Split(line, "%")

	re := regexp.MustCompile(`\r`)
	var windows bool = re.MatchString(line) //returns true if carriage return exists.
	var newSubString string
	if windows {
		newSubString = ".\n"
	} else {
		newSubString = ".\r\n"
	}
	var newLine string = lines[0]

	for i := 0; i < nrOfIndents; i++ { //adding the indentations
		newSubString += preString
	}
	if strings.Count(line, ". ") > 1 { //adding the newlines with indentations instead of fullstop + space.
		newLine = strings.Replace(newLine, ". ", newSubString, -1)
	}
	if len(lines) == 2 { //adding back the comment to the last newlined string.
		newLine += "%" + lines[1]
	}
	return newLine
}

//helper functions
