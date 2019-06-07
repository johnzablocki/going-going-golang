package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"

	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3"
)

//go get github.com/mattn/go-sqlite3
//go get -u github.com/jinzhu/gorm
//go get github.com/jmoiron/sqlx

//DBName file path
const DBName string = "./demo.db"

type (
	//BookAuthor struct
	BookAuthor struct {
		ID        int    `db:"id"`
		FirstName string `db:"firstName" gorm:"column:firstName"`
		LastName  string `db:"lastName" gorm:"column:lastName"`
	}
)

func getDB() *sql.DB {
	db, err := sql.Open("sqlite3", DBName) //convenient but not efficient

	Check(err)
	return db
}

func getDBX() *sqlx.DB {
	db, err := sqlx.Open("sqlite3", DBName) //convenient but not efficient

	Check(err)
	return db
}

func initDB() {
	os.Remove(DBName)
	db := getDB()
	defer db.Close()

	ddl := `CREATE TABLE book_authors (id integer not null primary key, firstName text, lastName text)`
	_, err := db.Exec(ddl)

	Check(err)
}

//Insert sample
func Insert() {
	initDB()

	db := getDB()
	defer db.Close()

	tx, err := db.Begin()
	Check(err)

	stmt, err := tx.Prepare("INSERT INTO book_authors VALUES (?, ?, ?)")
	Check(err)

	defer stmt.Close()

	_, err = stmt.Exec(1, "Hank", "Moody")
	Check(err)

	_, err = stmt.Exec(2, "Robert", "Ludlow")
	Check(err)

	_, err = stmt.Exec(3, "Tom", "Clancy")
	Check(err)

	tx.Commit()
}

//ReadAll reads all rows
func ReadAll() {
	Insert()

	db := getDB()
	rows, err := db.Query("SELECT id, firstName, lastName FROM book_authors")
	Check(err)
	defer rows.Close()

	for rows.Next() {
		var id int
		var firstName, lastName string

		err = rows.Scan(&id, &firstName, &lastName)
		Check(err)

		fmt.Printf("%d - %s, %s\n", id, lastName, firstName)
	}

}

//ReadRow reads a row
func ReadRow() {
	Insert()

	db := getDB()
	stmt, err := db.Prepare("SELECT id, firstName, lastName FROM book_authors WHERE id = ?")
	Check(err)
	defer stmt.Close()

	var id int
	var firstName, lastName string
	err = stmt.QueryRow(1).Scan(&id, &firstName, &lastName)
	Check(err)

	fmt.Printf("%d - %s, %s\n", id, lastName, firstName)
}

//ReadStruct reads a row and assigns to struct
func ReadStruct() {
	Insert()

	db := getDBX()
	var author BookAuthor
	err := db.Get(&author, "SELECT id, firstName, lastName FROM book_authors WHERE id = ?", 1)
	Check(err)

	printAuthor(author)
}

//ReadORM reads as an ORM
func ReadORM() {
	Insert()

	db, err := gorm.Open("sqlite3", DBName)
	Check(err)

	var author BookAuthor
	result := db.First(&author, "lastName = ?", "Moody")
	errs := result.GetErrors()
	for _, e := range errs {
		Check(e)
	}

	printAuthor(author)

}

func printAuthor(author BookAuthor) {
	fmt.Printf("%d - %s, %s\n", author.ID, author.LastName, author.FirstName)
}
