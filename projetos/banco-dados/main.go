package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	fmt.Println("Banco de Dados")

	urlBancoDados := "devbook:devbook@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, error := sql.Open("mysql", urlBancoDados)

	if error != nil {
		log.Fatal(error)
	}

	defer db.Close()

	if error = db.Ping(); error != nil {
		log.Fatal(error)
	}

	fmt.Println("Conectado com o banco de dados")

	records, error := db.Query("SELECT * FROM usuarios")

	if error != nil {
		log.Fatal(error)
	}

	defer records.Close()

	fmt.Println(records)
}
