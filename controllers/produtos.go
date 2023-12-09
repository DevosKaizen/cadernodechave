package controllers

import (
	"LOJAEMGO/models"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "Index", nil)

}
func Pegarchave(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "Pegarchave", nil)

}

// func Login(w http.ResponseWriter, r *http.Request) {

// 	temp.ExecuteTemplate(w, "Login", nil)

// }
func Salas(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosProdutos()
	temp.ExecuteTemplate(w, "Salas", todosOsProdutos)

}

func New(w http.ResponseWriter, r *http.Request) {

	// todosOsProdutos := models.BuscaTodosProdutos()
	// temp.ExecuteTemplate(w, "New", todosOsProdutos)

	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		username := r.FormValue("username")
		password := r.FormValue("password")

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço")
		}
		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão do quantidade")
		}

		errUsuario := models.CriarNovoUsuario(username, password)
		if errUsuario != nil {
			log.Println("Erro ao criar usuario:", errUsuario)
			http.Error(w, "Erro ao criar usuario. Tente novamente mais tarde.", http.StatusInternalServerError)
			log.Println("deu erro no controllers")

			return
		}

		models.CriarNovoProduto(nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt)
		models.CriarNovoUsuario(username, password)

	}
	http.Redirect(w, r, "/salas", 301)
}
func DeletaProduto(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Em controllers Passou por delete")
	idDoProduto := r.URL.Query().Get("id")

	// Verifica se o ID é um valor valido antes de continuar.
	if idDoProduto == "" {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		log.Println("Erro no delete")
		return
	}

	models.DeletaProduto(idDoProduto)
	http.Redirect(w, r, "/salas", 301)
	fmt.Println("Em models saiu do delete")
}
func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idDoProduto)
	temp.ExecuteTemplate(w, "Edit", produto)

}
func Update(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na converção do id para int: ", err)
		}
		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na converção do preço para float64: ", err)
		}
		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na converção da quantidade para int: ", err)
		}
		models.AtualizaProduto(idConvertidaParaInt, nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt)
	}
	http.Redirect(w, r, "/salas", 301)
}
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		user, err := models.GetUserByUsername(username)
		if err != nil {
			http.Error(w, "Usuario não encontrado", http.StatusUnauthorized)
			return
		}

		// Verificar senha - Lembre-se de usar um mecanismo de hash de senha na produção.
		if user.Password == password {
			// Autenticação bem-sucedida, pode redirecionar ou definir cookies/sessões.
			http.Redirect(w, r, "/salas", http.StatusFound)
			return
		}

		http.Error(w, "Credenciais invalidas", http.StatusUnauthorized)
		return
	}

	temp.ExecuteTemplate(w, "Login", 301)
}
func NewUser(w http.ResponseWriter, r *http.Request) {
	// Se a solicitacao for um POST, isso significa que o formulario foi enviado
	if r.Method == http.MethodPost {
		// Recupere os valores do formulario
		username := r.FormValue("username")
		password := r.FormValue("password")

		// CRIAR um novo usuario no banco de dados
		err := models.CriarNovoUsuario(username, password)
		if err != nil {
			http.Error(w, "Erro ao criar usuario", http.StatusInternalServerError)
			return
		}

		// Redirecione para a pagina de salas após a criacao bem-sucedida
		http.Redirect(w, r, "/users", http.StatusFound)
		return
	}

	// Se a solicitação não for um POST, exiba a pagina de criação de usuario
	temp.ExecuteTemplate(w, "NewUser", nil)
}

func Users(w http.ResponseWriter, r *http.Request) {

	selectDeTodosUsuarios := models.BuscaTodosUsuarios()

	temp.ExecuteTemplate(w, "Users", selectDeTodosUsuarios)

}
func DeletaUsuario(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Em controllers Passou por delete")
	idUsuario := r.URL.Query().Get("id")

	// Verifica se o ID é um valor valido antes de continuar.
	if idUsuario == "" {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		log.Println("Erro no delete")
		return
	}

	models.DeletaUsuario(idUsuario)
	http.Redirect(w, r, "/users", 301)
	fmt.Println("Em Controllers saiu do delete vai pra modells")
}

func Checkin(w http.ResponseWriter, r *http.Request) {
	proxSala := models.ProxSala()
	temp.ExecuteTemplate(w, "Checkin", proxSala)

}
