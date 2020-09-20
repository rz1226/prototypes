package main

import (
	"fmt"
	"github.com/rz1226/prototypes/filesread/files"
	"time"
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

func test_write() {
	for i := 0; i < 100; i++ {
		files.Write(i, "\n")
	}
	files.CloseWrite()
	time.Sleep(time.Second * 3)
}

func main() {
	test_write()
}
