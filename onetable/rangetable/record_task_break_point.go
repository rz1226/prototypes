package rangetable

import "github.com/rz1226/mysqlx"

const TABLE_NAME_TASK_RECORD_ID = "record_id"

type task struct {
	Id  int64 `orm:"id" auto:"1"`
	Id1 int64 `orm:"current_id"`
}

func getId1() (int64, error) {
	sql := "select current_id from " + TABLE_NAME_TASK_RECORD_ID + " where id = 1 limit 1 "
	res, err := mysqlx.SQLStr(sql).Query(Kit)
	if err != nil {
		return 0, err
	}
	return res.ToInt64()

}

func setId1(id1 int64) error {
	sql := "update " + TABLE_NAME_TASK_RECORD_ID + " set current_id = ? where id = 1 limit 1 "
	_, err := mysqlx.SQLStr(sql).AddParams(id1).Exec(Kit)
	return err

}

/*
create table  prototypes.record_id(
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `current_id` bigint  DEFAULT null COMMENT '任务记录点',

  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)

)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4  COMMENT='record_id';

insert into prototypes.record_id set current_id = 0 ,id = 1 ;

*/
