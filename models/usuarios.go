package models

import (
	"LOJAEMGO/db"
	"fmt"
)

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
	selectTodosUsuarios, err := db.Query("SELECT * FROM users")
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
