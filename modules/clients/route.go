package clients

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// registrando as rotas para esse modulo
func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/client", GetPeople).Methods("GET")
}

/**
*
* constroladores das rotas
*
 */

func GetPeople(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("rodando aqui")
}
