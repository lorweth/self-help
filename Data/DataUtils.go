package Data

import "io/ioutil"

func ReadFile(path string) []byte {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return data
}

func WriteFile(path string, data []byte) {
	err := ioutil.WriteFile(path, data, 0644)

	if err != nil {
		panic(err)
	} else {
		println("File written successfully")
	}
}
