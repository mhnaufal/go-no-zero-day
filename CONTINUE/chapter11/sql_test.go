package chapter11

import (
	"context"
	"fmt"
	"testing"
)

func TestInsertSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	qinsert := "INSERT INTO surat(nama) VALUES('utama');"
	_, err := db.ExecContext(ctx, qinsert)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully insert data")
}

func TestSelectSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	qselect := "SELECT id, nama FROM surat;"

	rows, err := db.QueryContext(ctx, qselect)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, nama string

		err = rows.Scan(&id, &nama)
		if err != nil {
			panic(err)
		}

		fmt.Printf("ID: %v\nNama: %v\n", id, nama)
	}
}

func TestQueryWithArguments(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	nama := "utama"
	qquery := "SELECT id, nama FROM surat WHERE nama = ? LIMIT 1"

	row, err := db.QueryContext(ctx, qquery, nama)

	if err != nil {
		panic(err)
	}
	defer row.Close()

	for row.Next() {
		var id, nama string

		err = row.Scan(&id, &nama)
		if err != nil {
			panic(err)
		}

		fmt.Printf("ID: %v\nNama: %v\n", id, nama)
	}
}

func TestInsertSqlAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	qinsert := "INSERT INTO surat(nama) VALUES('izin');"
	result, err := db.ExecContext(ctx, qinsert)
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully insert data with id ", insertId)
}
