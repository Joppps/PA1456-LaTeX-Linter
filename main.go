package main

import (
	File "LatexLinter/fileWrite"
	Lint "LatexLinter/pkg"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func isError(err error) bool {
	if err != nil {
		fmt.Println("Error!", err.Error())
	}
	return (err != nil)
}

func createSettings() {
	content := `{
	"NewLineFullStop": true,
	"IndentTabs": true,
	"IndentSpaces": 2,
	"Indentation": true,
	"FormatComments": true,
	"CommentTabs": false,
	"CommentSpacers": 4,
	"SectionBlanks": true,
	"SectionCount": 2,
	"ExcludedEnviroments":[
		"document"
	]
}`
	if err := os.WriteFile("settings.json", []byte(content), 0666); err != nil {
		log.Fatal(err)
	}
}

func getSettings() (Settings, error) {
	jsonFile, err := os.Open("settings.json")
	var settings Settings
	if isError(err) {
		fmt.Println("Settings.json not found! Creating new file :)")
		createSettings()
		jsonFile, err = os.Open("settings.json")
		if isError(err) {
			return settings, err
		}
	}
	defer jsonFile.Close()
	byteVal, err := ioutil.ReadAll(jsonFile)
	if isError(err) {
		return settings, err
	}
	json.Unmarshal(byteVal, &settings)
	return settings, err
}

type Settings struct {
	NewLineFullStop     bool
	IndentTabs          bool
	IndentSpaces        int
	Indentation         bool
	FormatComments      bool
	CommentTabs         bool
	CommentSpacers      int
	SectionBlanks       bool
	SectionCount        int
	ExcludedEnviroments []string
}

//returns true if extension is tex,bib or tikz
func checkFileExtension(fileName string) bool {
	var ret bool = false
	dots := strings.Split(fileName, ".")
	if dots[len(dots)-1] == "tex" || dots[len(dots)-1] == "bib" || dots[len(dots)-1] == "tikz" {
		ret = true
	}
	return ret
}

func main() {
	//Read settings file ---------------------------------------------
	settings, err := getSettings()
	if isError(err) {
		fmt.Println("Aborting!")
		return
	}
	//read TexFile -----------------------------------------------
	if !checkFileExtension(os.Args[1]) {
		fmt.Println("Wrong type of file! Must be tex, bib or tikz!")
		return
	}
	data, err := os.ReadFile(os.Args[1])
	if isError(err) {
		fmt.Println("Aborting!")
		return
	}
	lines := strings.Split(string(data), "\n") //blir array av typ string []string

	//Call Lint functions bellow-------------------------------------------------------------

	if settings.Indentation { //EnviromentIndentation kallar också (beroende på bool) på newLiner så den får korrekt indentation lite överallt.
		Lint.EnviromentIndentation(lines, len(lines), settings.IndentTabs, settings.IndentSpaces, settings.NewLineFullStop, settings.ExcludedEnviroments)
	} else if settings.NewLineFullStop {
		for i := 0; i < len(lines); i++ {
			lines[i] = Lint.NewLiner(lines[i], 0, "")
		}
	}
	if settings.FormatComments {
		for i := 0; i < len(lines); i++ {
			lines[i] = Lint.CommentFormat(lines[i], settings.CommentTabs, settings.CommentSpacers)
		}
	}
	if settings.SectionBlanks {
		for i := 0; i < len(lines); i++ {
			lines[i] = Lint.InsertBlankLinesOnSections(lines[i], settings.SectionCount)
		}
	}

	//Save or print output --------------------------------------------------------

	File.CreateAndWriteNewFile(os.Args[1], lines, "Linted")

}
