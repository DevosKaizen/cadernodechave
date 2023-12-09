package models

import (
	"LOJAEMGO/db"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosProdutos() []Produto {
	db := db.ConectaCombancoDeDados()
	selectDeTodosProdutos, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}
func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaCombancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade)values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}
func DeletaProduto(id string) {
	fmt.Println("Em models Passou por delete")
	db := db.ConectaCombancoDeDados()
	log.Println("Passou por conecta banco de dados")

	deletarOProduto, err := db.Prepare("delete from produtos where id=$1") //https://pkg.go.dev/database/sql#DB.Prepare <-- DOCUMENTAÇAO DO BANCO DE
	if err != nil {
		log.Println("deu erro no  entrou no if")
		panic(err.Error())
	}
	log.Println("saiu do if")
	deletarOProduto.Exec(id)
	log.Println("exec delete")
	defer db.Close()
	log.Println("fechou o servidor com defer")

}

func EditaProduto(id string) Produto {
	db := db.ConectaCombancoDeDados()

	produtoDoBanco, err := db.Query("select *from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}
	produtoParaAtualizar := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}
	defer db.Close()
	return produtoParaAtualizar

}
func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaCombancoDeDados()
	atualizaProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	atualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}

func ProxSala() Produto {
	db := db.ConectaCombancoDeDados()
	defer db.Close()

	// Seed é importante para garantir que você obtenha números diferentes em cada execução
	rand.Seed(time.Now().UnixNano())

	// Gera um número aleatório entre 0 e 10
	numeroAleatorio := rand.Intn(11) // Gera um número entre 0 e 10 (exclusivo)

	fmt.Println(numeroAleatorio)

	// Utiliza um prepared statement para evitar injeção de SQL
	proxSala, err := db.Prepare("SELECT id, nome, descricao, preco FROM produtos WHERE id = $1")
	if err != nil {
		return Produto{}
	}
	defer proxSala.Close()

	// Executa a query com o número aleatório como parâmetro
	produto, err := proxSala.Query(numeroAleatorio)
	if err != nil {
		return Produto{}
	}
	defer produto.Close()

	// Verifica se há linhas retornadas
	if produto.Next() {
		var p Produto
		err := produto.Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco)
		if err != nil {
			return Produto{}
		}
		return p
	}

	// Caso não haja linhas retornadas, você pode retornar um erro indicando isso
	return Produto{}
}

// Busca proxima sala

// func ProxSala() []Produto {
// 	db := db.ConectaCombancoDeDados() // verificar se nao tem comando random sql

// 	// Seed é importante para garantir que você obtenha números diferentes em cada execução
// 	rand.Seed(time.Now().UnixNano())

// 	// Gera um número aleatório entre 0 e 10
// 	numeroAleatorio := rand.Intn(11) // Gera um número entre 0 e 10 (exclusivo)

// 	fmt.Println(numeroAleatorio)

// 	proxSala, err := db.Query("select *from produtos where id=$1", numeroAleatorio)

// 	// proxSala, err := db.Query("SELECT id, nome, descricao, preco FROM produtos ORDER BY id ASC") BACKUP
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	p := Produto{}
// 	produtos := []Produto{}

// 	proxSala.Next()
// 	var id int
// 	var nome, descricao string
// 	var preco float64

// 	err = proxSala.Scan(&id, &nome, &descricao, &preco)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	p.Id = id
// 	p.Nome = nome
// 	p.Descricao = descricao
// 	p.Preco = preco

// 	produtos = append(produtos, p)

// 	defer db.Close()
// 	return produtos
// }
