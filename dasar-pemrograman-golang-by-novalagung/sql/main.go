package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	id    string
	name  string
	age   int
	grade int
}

func main() {
	// sqlQuery()
	// sqlQueryRow()
	// sqlPrepare()

	// A.56.6. Insert, Update, & Delete Data Menggunakan Exec()
	sqlExec()
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/db_belajar_golang", dbUser, dbPwd))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlQuery() {
	db, err := connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	age := 27
	rows, err := db.Query("SELECT id, name, grade FROM tb_student WHERE age = ?", age)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var result []student

	for rows.Next() {
		each := student{}
		err := rows.Scan(&each.id, &each.name, &each.grade)
		if err != nil {
			panic(err.Error())
		}

		result = append(result, each)
	}

	if err := rows.Err(); err != nil {
		panic(err.Error())
	}

	for _, each := range result {
		fmt.Println(each.name)
	}
}

func sqlQueryRow() {
	db, err := connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	result := student{}
	id := "E001"
	err = db.QueryRow("SELECT name, grade FROM tb_student WHERE id = ?", id).Scan(&result.name, &result.grade)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Name\t: %s\nGrade\t: %d\n", result.name, result.grade)
}

func sqlPrepare() {
	db, err := connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT name, grade FROM tb_student WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	result1 := student{}
	stmt.QueryRow("E001").Scan(&result1.name, &result1.grade)
	fmt.Printf("Name\t: %s\nGrade\t: %d\n", result1.name, result1.grade)

	result2 := student{}
	stmt.QueryRow("W001").Scan(&result2.name, &result2.grade)
	fmt.Printf("Name\t: %s\nGrade\t: %d\n", result2.name, result2.grade)

	result3 := student{}
	stmt.QueryRow("B001").Scan(&result3.name, &result3.grade)
	fmt.Printf("Name\t: %s\nGrade\t: %d\n", result3.name, result3.grade)
}

func sqlExec() {
	db, err := connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO tb_student VALUES (?, ?, ?, ?)", "G001", "Galahad", 29, 2)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Insert success!")

	_, err = db.Exec("UPDATE tb_student SET age = ? WHERE id = ?", 28, "G001")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("update success!")

	_, err = db.Exec("DELETE FROM tb_student WHERE id = ?", "G001")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("delete success!")
}
