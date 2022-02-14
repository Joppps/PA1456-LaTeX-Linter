package Lint

import "strings"

func NewLiner(line string, nrOfIndents int, spaces bool) string {
	var newLine string = line
	var newSubString string = ".\n"
	for i := 0; i < nrOfIndents; i++ {
		if spaces {
			newSubString += "    "
		} else {
			newSubString += "\t"
		}
	}
	if strings.Count(line, ". ") > 1 {
		newLine = strings.Replace(line, ". ", newSubString, -1)
	}
	return newLine
}

//helper functions
