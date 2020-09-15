package main

import (
	"fmt"
	"github.com/rz1226/prototypes/filesread/files"
)

func test_make_some_test_files() {
	files.MakeTestFiles()
}

func test_get_all_files() {
	path := "test_files/"

	res := files.GetAllFiles(path)

	for k, v := range res {
		fmt.Println(v)

		fmt.Println(files.GetFileContent(v))
		files.RemoveFile(v, fmt.Sprint(k)+".json")
	}
}

func main() {
	test_get_all_files()
}
