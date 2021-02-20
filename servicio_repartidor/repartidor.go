package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"../estructura"
)

var estadoRepartidor int =0
var hashPedido = make(map[int]estructura.PedidoRepartidor)

func recibir_pedido(w http.ResponseWriter, r *http.Request)  {

	data:= estructura.PedidoRepartidor{}
	m:= estructura.JSONMessageGeneric{}
	w.Header().Set("Content-Type","application/json")

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashPedido[data.IdPedido]=data

	m.Message =  "pedidio realizado"
	m.Id = data.IdPedido
	json.NewEncoder(w).Encode(m)
	fmt.Print("pedido recibida: ")
	fmt.Println(data)

}


func handle()  {

	router := mux.NewRouter()
	router.HandleFunc("/recibir_pedidio",recibir_pedido).Methods("POST")
	http.ListenAndServe(":8082", router)

}

func main (){

	handle()

}
