package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	fmt.Println("===================ENTROU")
	db, err := sql.Open("mysql", "root:debian23@/cursogo")
	if err != nil {
		log.Fatal(err)
		fmt.Println("=============FUDEU GO")
	}
	defer db.Close()

	rows, _ := db.Query("select id, nome from usuarios where id > ?", 3)
	defer rows.Close()

	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.nome)
		fmt.Println(u)
	}
}
