package rangetable

import (
	"fmt"
	"github.com/rz1226/mysqlx"
	"reflect"
)

var Kit *mysqlx.DB

func init() {
	mysqlx.Conf.TagName = "orm"
	f := func(tags reflect.StructTag) bool {
		tag := tags.Get("auto")
		if tag == "1" {
			return true
		}
		return false
	}
	mysqlx.Conf.FuncAuto = f
}

func init() {
	dbconf := mysqlx.NewDBConf(MYSQL_USERNAME, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_PORT, "prototypes", 12)
	kit, err := dbconf.Connect()

	if err != nil {
		fmt.Println(dbconf.Str(), err)
		panic("no db ")
	}
	Kit = kit

}
