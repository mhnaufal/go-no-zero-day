package chapter11

import (
	"context"
	"fmt"
	"strconv"
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

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	statement, err := db.PrepareContext(ctx, "INSERT INTO surat(nama) VALUES(?);")
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		result, err := statement.ExecContext(ctx, "surat "+strconv.Itoa(i))
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Insert surat ", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	qtransaction := "INSERT INTO surat(nama) VALUES(?);"

	for i := 0; i < 10; i++ {
		result, err := tx.ExecContext(ctx, qtransaction, "surat "+strconv.Itoa(i))
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Insert surat ", id)
	}

	err = tx.Rollback() // rollback transaction
	// err = tx.Commit()   // commit transaction
	if err != nil {
		panic(err)
	}
}
