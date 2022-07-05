package chapter11

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/mail_template_go")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
