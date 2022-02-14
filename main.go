package main

import (
	Lint "LatexLinter/pkg"
	"fmt"
	"os"
	"strings"
)

func isError(err error) bool {
	if err != nil {
		fmt.Println("Error!", err.Error())
	}
	return (err != nil)
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	if isError(err) {
		fmt.Println("Aborting!")
		return
	}
	lines := strings.Split(string(data), "\n") //blir array av typ string []string
	Lint.EnviromentIndentation(lines, len(lines))

	for i := 0; i < len(lines); i++ {
		fmt.Println(lines[i])
	}

}
