package routes

import (
	"LOJAEMGO/controllers"
	"net/http"
)

func CarregaRotas() {

	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/deletaproduto", controllers.DeletaProduto) // funçao modificada
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)
	http.HandleFunc("/salas", controllers.Salas)
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/pegarchave", controllers.Pegarchave)
	http.HandleFunc("/newuser", controllers.NewUser)
	http.HandleFunc("/users", controllers.Users)
	http.HandleFunc("/deletausuario", controllers.DeletaUsuario) // verificar essa funçao ok
	http.HandleFunc("/checkin", controllers.Checkin)
}
