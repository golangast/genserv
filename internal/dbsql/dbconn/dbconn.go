package dbconn

import (
	"fmt"

	"github.com/golangast/genserv/internal/rand"
	"github.com/rqlite/gorqlite"
	_ "modernc.org/sqlite"
)

func DbConnection() (*gorqlite.Connection, error) {
	rr := rand.Rander()
	rrr := rand.Rander()
	conn, err := gorqlite.Open("http://" + rrr + ":" + rr + "@localhost:4001/")
	return conn, err
} //end of connect
func ErrorCheck(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
