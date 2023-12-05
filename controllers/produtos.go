package controllers

import (
	"LOJAEMGO/models"
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

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço")
		}
		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão do quantidade")
		}
		models.CriarNovoProduto(nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt)

	}
	http.Redirect(w, r, "/", 301)
}
func Delete(w http.ResponseWriter, r *http.Request) {

	idDoProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idDoProduto)
	http.Redirect(w, r, "/", 301)

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
			log.Printf("Erro na converção do id para int: ", err)
		}
		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Printf("Erro na converção do preço para float64: ", err)
		}
		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Printf("Erro na converção da quantidade para int: ", err)
		}
		models.AtualizaProduto(idConvertidaParaInt, nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt)
	}
	http.Redirect(w, r, "/", 301)
}
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		user, err := models.GetUserByUsername(username)
		if err != nil {
			http.Error(w, "Usuário não encontrado", http.StatusUnauthorized)
			return
		}

		// Verificar senha - Lembre-se de usar um mecanismo de hash de senha na produção.
		if user.Password == password {
			// Autenticação bem-sucedida, pode redirecionar ou definir cookies/sessões.
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
		return
	}

	temp.ExecuteTemplate(w, "Login", 301)
}
