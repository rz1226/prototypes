package rangetable

import "github.com/rz1226/mysqlx"

const TABLE_NAME_2 = "check_and_insert"

//没有就插入

func CheckAndInsert(data string) error {
	sql := "insert into  " + TABLE_NAME_2 + " set file_name = ? "
	_, err := mysqlx.SQLStr(sql).AddParams(data).Exec(Kit)
	return err
}

/*


create table  prototypes.check_and_insert(
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `file_name` varchar(200)  DEFAULT null COMMENT '任务记录点',
  `status` int  not null default 0 ,
  `info` varchar(200) not null default '',
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  unique key(`file_name`)

)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4  COMMENT=' ';


*/
