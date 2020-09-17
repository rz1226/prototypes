package rangetable

import(
	"fmt"
	"github.com/rz1226/coroutinekit"
	"github.com/rz1226/mysqlx"
	"github.com/rz1226/utilx2/queue"

)
//把数据放入一个q，插入某表

const TABLE_NAME_INSERT = "example_insert"




var q *queue.BatchQueue
func init(){
	q = queue.NewBatchQueue(100,1)
	coroutinekit.Start("批量插库",1,fetchAndInsert, true, true )
}

type DataForInsert struct {
	ID     int64  `orm:"id" auto:"1"`
	Name   string `orm:"name"`
	Status int64  `orm:"status"`
}

func Put(data DataForInsert){
	q.Put(data )
}

func fetchAndInsert(){
	for {
		res, n , _ := q.Get(100)
		if n > 0 {
			data, ok  := res.([]interface{});if ok{
				dataInsert := make([]*DataForInsert, 0, len(data))
				for _, v := range data {
					v2 ,ok := v.(DataForInsert)
					if !ok{
						fmt.Println("fetchAndInsert not ok ")
					}
					dataInsert = append( dataInsert, &v2 )
				}
				sql, err := mysqlx.NewBM(&dataInsert).ToSQLInsert(TABLE_NAME_INSERT)
				if err != nil {
					fmt.Println( "get sql err :", err)

				}
				n, err := sql.Exec( Kit)
				fmt.Println("插入数据结果:",n)
				if err != nil {
					fmt.Println(sql.Info(), err)
				}
			}
		}

	}
}

/*

CREATE TABLE prototypes.example_insert (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20)  DEFAULT null COMMENT 'name',
   `status` int not null default 0 ,
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)

) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4  COMMENT='';


*/
