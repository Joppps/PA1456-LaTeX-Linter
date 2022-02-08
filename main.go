package main

import (
	"fmt"
	"os"
)

func isError(err error) bool {
	if err != nil {
		fmt.Println("Error!", err.Error())
	}
	return (err != nil)
}

func main() {
	fmt.Println("Opening file!")
	var file, err = os.OpenFile(".txtestt", os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()
}
