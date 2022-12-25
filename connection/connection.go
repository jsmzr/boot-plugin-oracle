package connection

import "database/sql"

var GlobalConn *sql.DB

func DB() *sql.DB {
	return GlobalConn
}
