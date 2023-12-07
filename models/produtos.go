package models

import (
	"LOJAEMGO/db"
	"fmt"
	"log"
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
	
	deletarOProduto, err := db.Prepare("delete from produtos where id=$1") //https://pkg.go.dev/database/sql#DB.Prepare
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

// USUARIOS

type User struct {
	ID       int
	Username string
	Password string
}

func GetUserByUsername(username string) (*User, error) {
	db := db.ConectaCombancoDeDados()
	defer db.Close()

	var user User
	err := db.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Password)
	return &user, err
}

// BuscaTodosUsuarios retorna todos os usuários cadastrados no banco de dados.
func BuscaTodosUsuarios() []User {
	db := db.ConectaCombancoDeDados()
	selectTodosUsuarios, err := db.Query("SELECT * FROM usuarios")
	if err != nil {
		panic(err.Error())
	}

	u := User{}
	usuarios := []User{}

	for selectTodosUsuarios.Next() {
		var id int
		var username, password string

		err = selectTodosUsuarios.Scan(&id, &username, &password)
		if err != nil {
			panic(err.Error())
		}
		u.ID = id
		u.Username = username
		u.Password = password
		usuarios = append(usuarios, u)
	}

	defer db.Close()
	return usuarios
}

// CriarNovoUsuario insere um novo usuário no banco de dados.
// func CriarNovoUsuario(username, password string) {
// 	db := db.ConectaCombancoDeDados()

// 	insereUsuario, err := db.Prepare("insert into usuarios(username, password) VALUES($1, $2)")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	insereUsuario.Exec(username, password)

//		defer db.Close()
//	}
func CriarNovoUsuario(username, password string) error {
	db := db.ConectaCombancoDeDados()

	insereUsuario, err := db.Prepare("INSERT INTO users(username, password) VALUES($1, $2)")
	if err != nil {
		fmt.Println("deu erro no models")
		return fmt.Errorf("erro ao preparar a consulta: %v", err)

	}

	_, err = insereUsuario.Exec(username, password)
	if err != nil {
		return fmt.Errorf("erro ao inserir usuário: %v", err)
	}

	defer db.Close()
	return nil
}
