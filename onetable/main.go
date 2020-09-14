package main

import (
	"fmt"
	"github.com/rz1226/prototypes/onetable/rangetable"
	"time"
)

func test_make_test_data() {
	err := rangetable.GenTestData()
	fmt.Println(err)
}

func test_range() {
	for {

		data := rangetable.GetNextData()
		time.Sleep(time.Millisecond * 100)
		fmt.Println(data)
		fmt.Println(rangetable.SetStatus(data.ID, 1))
	}
}

func test_check_insert() {
	for i := 0; i < 100; i++ {
		err := rangetable.CheckAndInsert("test_data_file_name_" + fmt.Sprint(i))
		fmt.Println(err)
	}
}

func main() {

	test_check_insert()

}
