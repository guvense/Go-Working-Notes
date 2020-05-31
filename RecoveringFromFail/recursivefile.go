package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	defer reportPanic()
	scanDirectory("go")
}

func reportPanic() {
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		panic(p)
	}
}

func scanDirectory(path string) error {

	fmt.Println(path)
	files, err := ioutil.ReadDir(path)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			scanDirectory(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
	return nil
}
