package main

import (
	"fmt"
	"strconv"
	"time"
)

func test(name string) {
	fmt.Println("my name is: ", name)
}

func isError(err error) bool {
	if err != nil {
		fmt.Println("Error!", err.Error())
	}
	return (err != nil)
}

func main() {
	// data, err := os.ReadFile(os.Args[1])
	// if isError(err) {
	// 	return
	// }
	// var test string = string(data)
	// lines := strings.Split(test, "\n") //blir array av typ string []string    <---- go
	// for i := 0; i < len(lines); i++ {
	// 	fmt.Println(lines[i])
	// }
	// fmt.Printf("type: %T\n", lines)

	for i := 0; i < 10; i++ {
		go test(strconv.Itoa(i))
	}
	time.Sleep(1 * time.Second)
}
