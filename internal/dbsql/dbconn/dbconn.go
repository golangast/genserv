package dbconn

import (
	"fmt"

	"github.com/rqlite/gorqlite"
	_ "modernc.org/sqlite"
)

func DbConnection() (*gorqlite.Connection, error) {
	conn, err := gorqlite.Open("http://bill:secret1@localhost:4001/")
	return conn, err
} //end of connect
func ErrorCheck(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
