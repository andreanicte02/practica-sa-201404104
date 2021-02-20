package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"../estructura"
)

var estadoRepartidor int =0 //uno ocupado 2 desocupado
var hashPedido = make(map[int]*estructura.PedidoRepartidor)

func recibir_pedido(w http.ResponseWriter, r *http.Request)  {

	data:= estructura.PedidoRepartidor{}

	m:= estructura.JSONMessageGeneric{}
	w.Header().Set("Content-Type","application/json")

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if estadoRepartidor == 0 {
		data.EstadoRepartidor = 1
	}
	hashPedido[data.IdPedido]=&data

	m.Message =  "pedidio realizado"
	m.Id = data.IdPedido
	json.NewEncoder(w).Encode(m)
	fmt.Print("pedido recibida: ")
	fmt.Println(data)

	if estadoRepartidor == 1{
		return
	}
	estadoRepartidor = 1
}


func informar_estado_cliente(w http.ResponseWriter, r *http.Request)  {

	data:= estructura.JSONGenerico{}

	m:= estructura.JSONMessageGeneric{}
	w.Header().Set("Content-Type","application/json")

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pedido,existePedido := hashPedido[data.Id]

	if !existePedido{
		mensaje_error,_ := json.Marshal(estructura.JSONMessageGeneric{"Ese pedido no existe",-1})
		http.Error(w, string(mensaje_error), http.StatusBadRequest)
		return
	}


	if pedido.EstadoRepartidor == -1{
		m.Message =" el pedido ya fue entreagado"
	} else if estadoRepartidor == 1 && pedido.EstadoRepartidor == 0 {
		m.Message="el reapartidor esta ocupado pero no con tu pedido"

	}else if estadoRepartidor == 1 && pedido.EstadoRepartidor == 1 {
		m.Message="el repartidor va en camino con tu pedido"
	} else if estadoRepartidor == 0 {
		m.Message="el reaprtidor no esta ocupado con ningun pedido"
	} else{
		m.Message="tu pedido ya fue entregado"
	}

	m.Id = data.Id
	json.NewEncoder(w).Encode(m)
	fmt.Print("pedido recibida: ")
	fmt.Println(data)

}

func marcar_pedido(w http.ResponseWriter, r *http.Request)  {

	data:= estructura.PedidoRepartidor{}
	m:= estructura.JSONMessageGeneric{}
	w.Header().Set("Content-Type","application/json")

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pedido,existePedido := hashPedido[data.IdPedido]

	if !existePedido{
		mensaje_error,_ := json.Marshal(estructura.JSONMessageGeneric{"Ese pedido no existe",-1})
		http.Error(w, string(mensaje_error), http.StatusBadRequest)
		return
	}


	estadoRepartidor = 0
	pedido.EstadoRepartidor = -1 //pedido entregado

	m.Id = data.IdPedido
	m.Message = "id pedido: "+ string(m.Id)+ " pedido completado"
	json.NewEncoder(w).Encode(m)
	fmt.Print("pedido recibida: ")
	fmt.Println(data)

}


func handle()  {

	router := mux.NewRouter()
	router.HandleFunc("/recibir_pedidio",recibir_pedido).Methods("POST")
	router.HandleFunc("/informar_estado_cliente",informar_estado_cliente).Methods("GET")
	router.HandleFunc("/marcar_pedido",marcar_pedido).Methods("POST")
	http.ListenAndServe(":8082", router)

}

func main (){

	handle()

}
