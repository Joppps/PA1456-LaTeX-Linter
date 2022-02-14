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
	NewLineFullStop bool
	IndentTabs      bool
	IndentSpaces    int
	Indentation     bool
	FormatComments  bool
	CommentTabs     bool
	CommentSpaces   int
}


func main() {
	//Read settings file ---------------------------------------------
	jsonFile, err := os.Open("settings.json")
	if isError(err) {
		fmt.Println("Settings.json not found!")
		//return???
	}
	defer jsonFile.Close()
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
	if settings.Indentation { //EnviromentIndentation kallar också (beroende på bool) på newLiner så den får korrekt indentation lite överallt.
		Lint.EnviromentIndentation(lines, len(lines), settings.IndentTabs, settings.IndentSpaces, settings.NewLineFullStop)
	} else if settings.NewLineFullStop {
		for i := 0; i < len(lines); i++ {
			lines[i] = Lint.NewLiner(lines[i], 0, "")
		}
	}

	//Save or print output --------------------------------------------------------

	for i := 0; i < len(lines); i++ {
		fmt.Println(lines[i])
	}

	//Print to check JSON
	//fmt.Println("NewLineFullStop:", settings.NewLineFullStop)
	//fmt.Println("IndentSettings:", "\n\tTabs:", settings.IndentTabs)
	//fmt.Println("\tSpaces:", settings.IndentSpaces)
	//fmt.Println("\tIndentation:", settings.Indentation)
	//fmt.Println("CommentSettings:")
	//fmt.Println("\tFormatComments:", settings.FormatComments)
	//fmt.Println("\tCommentTabs:", settings.CommentTabs)
	//fmt.Println("\tCommentSpaces:", settings.CommentSpaces)
}
