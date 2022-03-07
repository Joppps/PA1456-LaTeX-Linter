package File

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//will make a filecopy of the file pointed to by filepath, filename: filename
func CreateAndWriteNewFile(filepath string, lines []string, fileNameExtender string) bool {
	var success bool = true

	temp := strings.Split(filepath, string(os.PathSeparator))
	fileTotalName := temp[len(temp)-1]
	fileParts := strings.Split(fileTotalName, ".") //please dont have dots in your fileNames....

	if strings.Count(fileTotalName, ".") > 1 {
		success = false
		fmt.Println("No files with dots in filename allowed!")
		return success
	}

	//create fileCopyName
	var fileNameGood bool = false
	var fileNameTest string
	for i := 0; i < 30 && !fileNameGood; i++ { //you really shouldnt have more than 30 of lintedFiles lol...
		if i == 0 {
			fileNameTest = fileParts[0] + fileNameExtender + "." + fileParts[1] //Name + Linter + . + tex
		} else {
			fileNameTest = fileParts[0] + fileNameExtender + strconv.Itoa(i) + "." + fileParts[1] //adds the number after fileName Extender
		}
		fileNameGood = !fileExist(fileNameTest)
	}
	if fileNameGood {
		file, err := os.Create(fileNameTest)
		if err != nil {
			fmt.Println("Error!", err)
			success = false
			return success
		}
		defer file.Close()
		for _, line := range lines {
			if strings.Contains(line, "\r") { //windows line endings uses \r\n hopefully all lines are the same ....
				_, err := file.WriteString(line + "\r\n")
				if err != nil {
					log.Fatal(err)
				}
			} else {
				_, err := file.WriteString(line + "\n")
				if err != nil {
					log.Fatal(err)
				}
			}

		}
	}
	fmt.Println("Successfully written to: ", fileNameTest)
	return success
}

func fileExist(fileName string) bool { //returns true if it exists
	_, err := os.Stat(fileName)
	return !errors.Is(err, os.ErrNotExist)
}
