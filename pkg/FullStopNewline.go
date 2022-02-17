package Lint

import (
	"strings"
)

func NewLiner(line string, nrOfIndents int, preString string) string {
	//split the string into two and then make the fullstop thingy on the non comment part.
	//then add the comment back onto the newLine added string.

	lines := strings.Split(line, "%")

	var newSubString string = ".\n"
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
