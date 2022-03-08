package Lint

import (
	"regexp"
	"strings"
)

//split the string into two and then make the fullstop thingy on the non comment part.
//then add the comment back onto the newLine added string.
func NewLiner(line string, nrOfIndents int, preString string) string {

	lines := strings.Split(line, " %")
	if len(lines) == 1 {
		lines = strings.Split(line, "%")
	}

	re := regexp.MustCompile(`\r`)
	var windows bool = re.MatchString(line) //returns true if carriage return exists.
	var newSubString string
	if windows {
		newSubString = "\r\n"
	} else {
		newSubString = "\n"
	}
	var newLine string = lines[0] //divides the comment part and the latex part. latex part is now iun newLine.
	re = regexp.MustCompile(`(?m)((\. )|(\.)([A-z]))`)
	for i := 0; i < nrOfIndents; i++ { //adding the indentations
		newSubString += preString
	}
	newLine = re.ReplaceAllString(newLine, "$2$3"+newSubString+"$4")

	// if strings.Count(line, ". ") > 1 { //adding the newlines with indentations instead of fullstop + space.
	// 	newLine = strings.Replace(newLine, ". ", newSubString, -1)
	// }

	if len(lines) >= 2 {
		newLine += " %" + lines[1]
		for i := 2; i < len(lines); i++ {
			newLine += lines[i]
		}
	}
	return newLine
}

//helper functions
