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

func test_batch_insert(){
	for i := 0; i < 10000; i++ {
		data := rangetable.DataForInsert{}
		data.Name = fmt.Sprint("name_",i )
		rangetable.Put(data)

	}


	time.Sleep(time.Second*5)
}

func main() {
	for i := 0; i < 100; i++ {
		err := rangetable.Update("example_insert", fmt.Sprint("id = ",i),  map[string]interface{}{"status":1042})
		fmt.Println(err )
	}
}
