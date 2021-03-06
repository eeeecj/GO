package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:pNFHVsRw@tcp(39.101.135.40:3306)/gobasic")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var (
		id   int
		name string
	)

	rows, err := db.Query("select id,name from users where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
