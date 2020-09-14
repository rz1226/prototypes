package rangetable

import (
	"os"
	"strings"
)

var MYSQL_HOST = strings.TrimSpace(os.Getenv("MYSQL_HOST"))
var MYSQL_PORT = strings.TrimSpace(os.Getenv("MYSQL_PORT"))
var MYSQL_USERNAME = strings.TrimSpace(os.Getenv("MYSQL_USERNAME"))
var MYSQL_PASSWORD = strings.TrimSpace(os.Getenv("MYSQL_PASSWORD"))
