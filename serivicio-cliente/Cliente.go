package main

import (
	"./utils"
	"./peticiones-cliente"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
)




func menuRandom()  int {

	return rand.Intn(2 - 0) + 0
}


func solicitarPedido(w http.ResponseWriter, r *http.Request)  {


	data:= utils.JSONGenerico{}

	m:= utils.JSONMessageGeneric{}
	w.Header().Set("Content-Type","application/json")

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	m=peticiones_cliente.PeticionSolicitarPedido(&utils.Pedido{menuRandom(),data.Id,0})
	peticiones_cliente.IdPedido = m.Id
	json.NewEncoder(w).Encode(m)
	utils.LogSalida(data,m)

}

func getEstadoRestaurante(w http.ResponseWriter, r *http.Request)  {


	data:= utils.JSONGenerico{}

	m:= utils.JSONMessageGeneric{}
	w.Header().Set("Content-Type","application/json")

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	m= peticiones_cliente.PeticionSolicitarEstadoRestaurante(&data)
	json.NewEncoder(w).Encode(m)

	utils.LogSalida(data,m)



}

func getEstadoRepartidor(w http.ResponseWriter, r *http.Request)  {


	data:= utils.JSONGenerico{}

	m:= utils.JSONMessageGeneric{}
	w.Header().Set("Content-Type","application/json")

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	m= peticiones_cliente.PeticionEstadoRepartidor(&data)
	json.NewEncoder(w).Encode(m)
	utils.LogSalida(data,m)




}


func handle() {

	router := mux.NewRouter()
	router.HandleFunc("/solicitar_pedido",solicitarPedido).Methods("POST")
	router.HandleFunc("/get_estado_restaurante",getEstadoRestaurante).Methods("GET")
	router.HandleFunc("/get_estado_repartidor",getEstadoRepartidor).Methods("GET")
	http.ListenAndServe(":8080", router)
}





func main()  {

	utils.RegistrarServicio(&utils.ServicioData{"8080", "solicitar_pedido","/solicitar_pedido","cliente","POST"}, "POST","8085","/registrar_microservicio")
	utils.RegistrarServicio(&utils.ServicioData{"8080", "get_estado_restaurante","/get_estado_restaurante","cliente","GET"}, "POST","8085","/registrar_microservicio")
	utils.RegistrarServicio(&utils.ServicioData{"8080", "get_estado_repartidor","/get_estado_repartidor","cliente","GET"}, "POST","8085","/registrar_microservicio")

	fmt.Println("Escuhando puerto 8080")
	peticiones_cliente.HashPedido = make(map[int]int)
	peticiones_cliente.Codigo  = -1
	peticiones_cliente.IdPedido = -1
	handle()


}