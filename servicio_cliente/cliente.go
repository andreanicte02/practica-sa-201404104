
package main

import (
	"../estructura"
	"./peticiones_cliente"
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
)




func menuRandom()  int {

	return rand.Intn(2 - 0) + 0
}


func solicitar_pedido(w http.ResponseWriter, r *http.Request)  {


	data:= estructura.JSONGenerico{}

	m:= estructura.JSONMessageGeneric{}
	w.Header().Set("Content-Type","application/json")

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	m= peticiones_cliente.Peticion_solicitar_pedido(&estructura.Pedido{menuRandom(),data.Id,0})
	peticiones_cliente.IdPedido = m.Id
	json.NewEncoder(w).Encode(m)
	estructura.LogSalida(data,m)

}

func get_estado_restaurante(w http.ResponseWriter, r *http.Request)  {


	data:= estructura.JSONGenerico{}

	m:= estructura.JSONMessageGeneric{}
	w.Header().Set("Content-Type","application/json")

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	m= peticiones_cliente.Peticion_solicitar_estado_restaurante(&data)
	json.NewEncoder(w).Encode(m)

	estructura.LogSalida(data,m)



}

func get_estado_repartidor(w http.ResponseWriter, r *http.Request)  {


	data:= estructura.JSONGenerico{}

	m:= estructura.JSONMessageGeneric{}
	w.Header().Set("Content-Type","application/json")

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	m= peticiones_cliente.Peticion_estado_repartidor(&data)
	json.NewEncoder(w).Encode(m)
	estructura.LogSalida(data,m)




}


func handle() {

	router := mux.NewRouter()
	router.HandleFunc("/solicitar_pedido",solicitar_pedido).Methods("POST")
	router.HandleFunc("/get_estado_restaurante",get_estado_restaurante).Methods("GET")
	router.HandleFunc("/get_estado_repartidor",get_estado_repartidor).Methods("GET")
	http.ListenAndServe(":8080", router)
}



func main()  {

	 peticiones_cliente.HashPedido = make(map[int]int)
	 peticiones_cliente.Codigo  = -1
	 peticiones_cliente.IdPedido = -1
	 handle()


}
