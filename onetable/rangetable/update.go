package rangetable

import (
	"github.com/rz1226/mysqlx"
	"strings"

)

func Update(table string,  where string , data map[string]interface{} ) error {

	sqlstr :=  "update "+ table + " set "
	parmas := make([]interface{},0,len(data))
	for k,v := range data {
		sqlstr += k + " = ?,"
		parmas = append( parmas, v )
	}

	sql := mysqlx.SQLStr( strings.TrimRight(sqlstr,",") + " where "+ where ).AddParams(parmas...)
	//fmt.Println(sql.Info() )

	_, err := sql.Exec(Kit )
	return err
}
