package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123@QWE@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?,?)")
	stmt.Exec(2000, "Bia")
	stmt.Exec(2001, "Carlos")
	_, err1 := stmt.Exec(1, "Tiago") // Chave duplicada

	if err1 != nil {
		tx.Rollback()
		log.Fatal(err1)
	}

	tx.Commit()
}
