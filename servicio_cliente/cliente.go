
package main

import (
	"../estructura"
	"github.com/gorilla/mux"
	"net/http"
)


var hash = make(map[int]estructura.Cliente)

func solicitar_pedido(w http.ResponseWriter, r *http.Request){



}

func handle() {

	router := mux.NewRouter()
	router.HandleFunc("/crear_pedido", solicitar_pedido).Methods("POST")
	http.ListenAndServe(":8080", router)
}


func main()  {

	hash[0]= estructura.Cliente{1,"cliente1"}
	hash[0]= estructura.Cliente{2,"cliente2"}
	hash[0]= estructura.Cliente{3,"cliente3"}




}
