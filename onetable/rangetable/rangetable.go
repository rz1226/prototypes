package rangetable

import (
	"fmt"
	"github.com/rz1226/coroutinekit"
	"github.com/rz1226/mysqlx"
	"time"
)

const BATCH_SIZE = 10
const TABLE_NAME_RANGE = "example_range"

var c chan ExampleRange

func init() {
	c = make(chan ExampleRange, 0)
	coroutinekit.Start("test", 1, getNextDataTask, true, true)
}

func GetNextData() ExampleRange {
	data := <-c
	return data
}

func SetStatus(id int64, status int64) error {
	sql := "update " + TABLE_NAME_RANGE + " set status = ? where id = ? limit 1 "
	_, err := mysqlx.SQLStr(sql).AddParams(status, id).Exec(Kit)
	return err

}

func getNextDataTask() {
	for {
		currentId, err := getId1()
		fmt.Println("currentid =", currentId)
		if err != nil {
			fmt.Println(err)
		}

		sql := "select id, name from " + TABLE_NAME_RANGE + " where id >=  ? and status = 0  order by id asc limit  " + fmt.Sprint(BATCH_SIZE)
		res, err := mysqlx.SQLStr(sql).AddParams(currentId).Query(Kit)
		if err != nil {
			fmt.Println("err", err)
			return
		}

		var datas []*ExampleRange
		err = res.ToStruct(&datas)
		if err != nil {
			fmt.Println("任务跑完了 :", err)
			time.Sleep(time.Second * 12)
			return
		}

		for _, v := range datas {
			c <- *v
			currentId = v.ID

		}

		err = setId1(currentId + 1)
		if err != nil {
			fmt.Println(err)
		}
		if len(datas) < BATCH_SIZE {
			time.Sleep(time.Second * 10)
			continue
		}

	}

}

type ExampleRange struct {
	ID     int64  `orm:"id" auto:"1"`
	Name   string `orm:"name"`
	Status int64  `orm:"status"`
}

/*

CREATE TABLE prototypes.example_range (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20)  DEFAULT null COMMENT 'name',
   `status` int not null default 0 ,
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)

) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4  COMMENT='';


*/
