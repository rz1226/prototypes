package rangetable

import (
	"fmt"
	"github.com/rz1226/mysqlx"
)

const TABLE_NAME = "example_range"

type Data struct {
	ID   int64  `orm:"id" auto:"1"`
	Name string `orm:"name"`
}

func GenTestData() error {
	res := make([]*Data, 0)

	for i := 0; i < 1000; i++ {
		data := &Data{}
		data.Name = "name_" + fmt.Sprint(i)
		res = append(res, data)

	}
	sql, err := mysqlx.NewBM(&res).ToSQLInsert(TABLE_NAME)
	if err != nil {
		return err
	}
	n, err := sql.Exec(Kit)
	if err != nil {
		return err
	}

	fmt.Println("已经插入=", n)
	return nil
}
