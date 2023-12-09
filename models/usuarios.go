package models

import (
	"LOJAEMGO/db"
	"database/sql"
	"fmt"
	"log"
)

// USUARIOS

type User struct {
	Id       int
	Username string
	Password string
}

func GetUserByUsername(username string) (*User, error) {
	db := db.ConectaCombancoDeDados()
	defer db.Close()
	// TENTEI UTILIZAR NEW COMO E NAO TEM O COMPORTAMENTO ESPERADO
	var user User
	err := db.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username).Scan(&user.Id, &user.Username, &user.Password)
	return &user, err
}

// BuscaTodosUsuarios retorna todos os usuários cadastrados no banco de dados.
func BuscaTodosUsuarios() []User {
	db := db.ConectaCombancoDeDados()
	selectTodosUsuarios, err := db.Query("select * from users order by id asc")
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
		u.Id = id
		u.Username = username
		u.Password = password
		usuarios = append(usuarios, u)
	}

	defer db.Close()
	return usuarios
}

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
func DeletaUsuario(id string) {
	fmt.Println("Chegou em models Passou por func DeletaUsuario(id string)")
	db := db.ConectaCombancoDeDados()
	log.Println("Passou por db.ConectaCombancoDeDados()")

	deletarOUsuario, err := db.Prepare("delete from users where id=$1") //https://pkg.go.dev/database/sql#DB.Prepare <-- DOCUMENTAÇAO DO BANCO DE DADOS
	if err != nil {
		log.Println("deu erro no  entrou no if")
		panic(err.Error())
	}
	log.Println("saiu do if")
	deletarOUsuario.Exec(id)
	log.Println("executou delete")
	defer db.Close()
	log.Println("fechou o servidor com defer, retorna a users")

}

func GetNextProductID(db *sql.DB) (int, error) {
	var nextID int
	err := db.QueryRow("SELECT MAX(id) FROM produtos").Scan(&nextID)
	if err != nil {
		return 0, err
	}
	// Incrementar para obter o próximo ID
	nextID++
	return nextID, nil
}
