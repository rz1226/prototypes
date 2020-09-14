package files

import (
	"fmt"
	"os"
	"strings"
)

//制造一些测试用的文件

func MakeTestFiles() error {
	path := "test_files/"

	for i := 10000; i < 20000; i++ {
		fileName := "test_file_" + fmt.Sprint(i) + ".json"
		dir := fmt.Sprint(i/100) + "/"
		dirAll := path + dir
		err := CreateFile(dirAll, fileName, "content no is "+fmt.Sprint(i))
		if err != nil {
			fmt.Println("创建文件出错:", err)
			return err
		}
		fmt.Println("创建文件成功", dirAll, fileName)
	}

	return nil
}

//会覆盖已经存在的文件
func CreateFile(path string, fileName string, content string) error {
	res, err := PathExists(path)
	if err != nil {
		return err
	}
	if !res {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	newFile, err := os.Create(strings.TrimRight(path, "/") + "/" + fileName)
	if err != nil {
		return err
	}

	if content != "" {
		newFile.WriteString(content)
	}
	defer newFile.Close()
	return nil
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
