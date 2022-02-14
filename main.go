package main

import (
	Lint "LatexLinter/pkg"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func isError(err error) bool {
	if err != nil {
		fmt.Println("Error!", err.Error())
	}
	return (err != nil)
}

type Settings struct {
}

func main() {
	//Read settings file ---------------------------------------------
	jsonFile, err := os.Open("settings.json")
	if isError(err) {
		fmt.Println("Settings.json not found!")
		//return???
	}
	var settings Settings
	byteVal, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteVal, &settings)

	//read TexFile -----------------------------------------------
	data, err := os.ReadFile(os.Args[1])
	if isError(err) {
		fmt.Println("Aborting!")
		return
	}
	lines := strings.Split(string(data), "\n") //blir array av typ string []string

	//Call Lint functions bellow-------------------------------------------------------------

	Lint.EnviromentIndentation(lines, len(lines)) //indenterar och lägger newlines.

	//Svave or print output --------------------------------------------------------

	for i := 0; i < len(lines); i++ {
		fmt.Println(lines[i])
	}

}
