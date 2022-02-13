package indentation

import "strings"

//Inserts tabs depending on all lines in the strings depending on the number of nested enviroments
func EnviromentIndentation(fileLines []string, nrOfLines int) {
	var enviroments uint8 = 0
	for i := 0; i < nrOfLines; i++ {
		if ender(fileLines[i]) {
			enviroments--
		}

		var newLine string
		for i := 0; i < int(enviroments); i++ {
			newLine += "\t"
		}
		newLine += fileLines[i]
		fileLines[i] = newLine

		if starter(fileLines[i]) {
			enviroments++
		}
	}
}

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
