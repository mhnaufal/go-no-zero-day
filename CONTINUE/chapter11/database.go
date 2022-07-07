package chapter11

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/mail_template_go2?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
