package Lint

import (
	"regexp"
	"strings"
)

//CapitalLetter on first character for function gets included to other packages I.E Public usage.

//Inserts tabs depending on all lines in the strings depending on the number of nested enviroments
func EnviromentIndentation(fileLines []string, nrOfLines int, tabs bool, nrOfSpaces int, fullStopNewLine bool) {
	var enviroments int = 0
	var preString string = "\t"
	if !tabs { //tabs = false => spaces = true
		preString = ""
		for i := 0; i < nrOfSpaces; i++ {
			preString += " "
		}
	}

	for i := 0; i < nrOfLines; i++ {

		if ender(fileLines[i]) {
			enviroments--
		}
		var newLine string
		for i := 0; i < enviroments; i++ {
			newLine += preString
		}
		var removedWhiteSpaceLine = removeIndent(fileLines[i])
		newLine += removedWhiteSpaceLine
		fileLines[i] = newLine

		if starter(fileLines[i]) {
			enviroments++
		}
		if fullStopNewLine {
			fileLines[i] = NewLiner(newLine, enviroments, preString)
		}
	}
}

func removeIndent(line string) string {
	re := regexp.MustCompile(`(?m)^\s*`) //match all whitespace from start, multiline = true.
	line = re.ReplaceAllString(line, "") //removing all whitespace from start.
	return line
}

//helper functions
func starter(line string) bool {
	var contains bool = false
	begin, end := "\\begin{", "\\end{"
	docBegin := "\\begin{document}"
	if strings.Contains(line, docBegin) {
		return contains
	}
	if strings.Contains(line, begin) && !strings.Contains(line, end) {
		contains = true
	}
	return contains
}

func ender(line string) bool {
	var contains bool = false
	begin, end := "\\begin{", "\\end{"
	docEnd := "\\end{document}"
	if strings.Contains(line, docEnd) {
		return contains
	}
	if strings.Contains(line, end) && !strings.Contains(line, begin) {
		contains = true
	}
	return contains
}
