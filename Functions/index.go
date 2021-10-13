package Functions

import (
	"io/ioutil"
	"strings"
)

func CreateMutilpleInsertStringFromFile(readpath, writepath string) error {
	data, err := ioutil.ReadFile(readpath)

	if err != nil {
		return err
	}

	str := string(data)

	strArr := strings.Split(str, "\n")

	tableName := strArr[0]
	columns := strings.Split(strArr[1], ",")
	var values [][]string

	for _, v := range strArr[2:] {
		values = append(values, strings.Split(v, ","))
	}

	insertStrs := CreateMutilpleInsertString(tableName, columns, values)

	var output []byte
	for _, v := range insertStrs {
		output = append(output, []byte(v+"\n")...)
	}

	err = ioutil.WriteFile(writepath, output, 0644)
	if err != nil {
		return err
	}

	return nil
}
