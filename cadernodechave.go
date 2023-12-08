package main

import (
	"LOJAEMGO/routes"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	// //nao vamos mais usar isto
	// db := conectaCombancoDeDados()
	// defer db.Close()
	routes.CarregaRotas()

	http.ListenAndServe("192.168.0.15:8000", nil)

}
