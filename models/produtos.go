package models

import "guilherme/estudando-go/db"

type Product struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Product {
	db := db.ConectaComBanco()

	selectdetodososprodutos, err := db.Query("select * from product")
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
	defer db.Close()
	return produtos
}
