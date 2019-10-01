package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"streamViewPro/api/defs"
)


var (
	dbConn *sql.DB
	err error
)

func init() {
	db, e := sql.Open(
		"mysql", "thesevensky:xy3055789@tcp(localhost:3306)/videoPro?charset=utf8")
	defs.CheckErrorForExitOfMsg(e, "数据库连接错误")
	dbConn = db
}
