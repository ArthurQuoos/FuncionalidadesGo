package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConexao := "root: (SENHA) @/CursoGo?charset=utf8&parseTime=true&loc=Local" // tenho quue trocar pelo meu bd
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")
	linhas, erro := db.Query("SELECT * FROM usuarios")

	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()
	fmt.Println(linhas)

}
