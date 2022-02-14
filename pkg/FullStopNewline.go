package Lint

import "strings"

func NewLiner(line string, nrOfIndents int, preString string) string {
	var newLine string = line
	var newSubString string = ".\n"
	for i := 0; i < nrOfIndents; i++ {
		newSubString += preString
	}
	if strings.Count(line, ". ") > 1 {
		newLine = strings.Replace(line, ". ", newSubString, -1)
	}
	return newLine
}

//helper functions
