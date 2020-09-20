package files

import (
	"fmt"
	"github.com/rz1226/coroutinekit"
	"os"
)

var c chan string

func init() {
	c = make(chan string, 10000)
	coroutinekit.Start("写入文件", 1, writeToFile, true, false)

}
func Write(v ...interface{}) {
	str := fmt.Sprint(v...)
	c <- str
}
func CloseWrite() {
	close(c)
}

func writeToFile() {
	filename := "xx.txt"

	newFile, err := os.Create(filename)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer newFile.Close()

	for v := range c {
		newFile.WriteString(v)
	}

}
