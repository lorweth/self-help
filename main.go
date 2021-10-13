package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"io.witcher.self-help/Data"
	"io.witcher.self-help/Functions"
)

func makeInsertString(readpath, writepath string) {
	content := string(Data.ReadFile(readpath))
	var contentArr = strings.Split(content, "\n")

	tablename := contentArr[0]
	columns := strings.Split(contentArr[1], ",")
	values := strings.Split(contentArr[2], ",")

	var insertString = Functions.MakeInsertQueryString(tablename, columns, values)
	Data.WriteFile(writepath, []byte(insertString))
}

func makeMutipleInsertString(readpath, writepath string) {
	content := string(Data.ReadFile(readpath))
	var contentArr = strings.Split(content, "\n")

	tablename := contentArr[0]
	columns := strings.Split(contentArr[1], ",")
	var values [][]string

	for _, value := range contentArr[2:] {
		values = append(values, strings.Split(value, ","))
	}

	var insertString = Functions.MakeMultipleInsertQueryString(tablename, columns, values)
	var output []byte
	for _, content := range insertString {
		output = append(output, content...)
		output = append(output, '\n')
	}
	Data.WriteFile(writepath, output)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	line := scanner.Text()

	switch strings.ToLower(line) {
	case "go":
		break
	case "makeinsert":
		fmt.Println("Enter path to the input...")
		scanner.Scan()
		readpath := scanner.Text()
		fmt.Println("Enter path to the output...")
		scanner.Scan()
		writepath := scanner.Text()
		makeInsertString(readpath, writepath)
		break
	case "makemutilpleinsert":
		fmt.Println("Enter path to the input...")
		scanner.Scan()
		readpath := scanner.Text()
		fmt.Println("Enter path to the output...")
		scanner.Scan()
		writepath := scanner.Text()
		makeMutipleInsertString(readpath, writepath)
		break
	}

	// data := ""

	// for scanner.Scan() {
	// 	if strings.ToLower(scanner.Text()) == "go" {
	// 		content := string(Data.ReadFile(data))
	// 		var contentArr = strings.Split(content, "\n")
	// 		for _, value := range contentArr {
	// 			fmt.Println(value)
	// 		}
	// 	} else {
	// 		data = data + scanner.Text()
	// 	}
	// }

}
