package day_28_29

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type student struct {
	id    string
	name  string
	age   int
	grade int
}

func Sql() {
	sqlQuery()
	sqlPrepare()
}

func connect() (*sql.DB, error) {
	// Open up connection to the database
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_belajar_golang")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlQuery() {
	/// create a connection
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	// var age int = 33

	/// Execute the query
	/// return a instance sql.*Rows
	/// db.Query will result multiple rows
	// rows, err := db.Query("SELECT id, name, grade FROM tb_student WHERE age = ?", age)
	rows, err := db.Query("SELECT * FROM tb_student")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []student

	for rows.Next() {
		each := student{}

		/// Copy the columns value from 'rows' into variable 'each'
		/// 'rows' must match with the 'each' variable
		var err = rows.Scan(&each.id, &each.name, &each.age, &each.grade)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Printf("%v -- %v\n", each.name, each.age)
	}
}

func sqlPrepare() {
	/// create a connection to the database
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	/// prepare the query
	query, err := db.Prepare("SELECT name, grade FROM tb_student WHERE id ?")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	/// create one student object
	var result1 = student{}

	/// run the prepared query
	query.QueryRow("E001").Scan(&result1.name, &result1.grade)
	fmt.Printf("Name: %v -- Grade %v\n", result1.name, result1.grade)
}
