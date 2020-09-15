package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

//把一个目录中所有的流水账单找出来
func GetAllFiles(path string) []string {
	res := make([]string, 0)

	f := func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if strings.Contains(p, ".json") {
			res = append(res, p)
		}
		return nil
	}

	err := filepath.Walk(path, f)
	if err != nil {

		return nil
	}
	return res

}

//读取文件内容
func GetFileContent(file string) (string, error) {
	res, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("读取文件内容方法失败", err)
		return "", err
	}
	return string(res), nil
}
