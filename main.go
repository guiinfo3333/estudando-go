package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

func conectaComBanco() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=admin host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Nao deu erro")
	}
	return db
}

type Product struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectaComBanco()

	selectdetodososprodutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}

	produtos := []Product{}
	for selectdetodososprodutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectdetodososprodutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	// produtos := []Product{
	// 	{Nome: "Camiseta", Descricao: "Azul bonito", Preco: 39, Quantidade: 2},
	// 	{"Tenis", "Confortável", 89, 3},
	// 	{"Fone", "Muito Bom", 59, 2},
	// }
	temp.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}
