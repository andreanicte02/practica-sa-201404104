package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"../estructura"
)



var estadoRepartidor int =0 //-1 indica que no tiene pedido pendiente, 0 que esta en caminio y 1 que ya se entrego
//se isa una hash en el lado del reapartidor para llevarla en memoria
var hashPedido = make(map[int]*estructura.PedidoRepartidor)

/**
	w -> indica el contenido de la respuesta
	r -> indica el contenido de la solicitud
*/


//servicio que recibe el pedido, indica al reapartidor que tiene pedidos que entregar
//recibe una json con la estrucutra de PedidoRepartidor y actualiza el estado
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

//Recibe un json generico, con el id del pedido
//Rregresa el estado del pedido
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
		mensaje_error,_ := json.Marshal(estructura.JSONMessageGeneric{"El repartidor sigue en espera del pedido",-1})
		http.Error(w, string(mensaje_error), http.StatusBadRequest)
		fmt.Print("pedido recibida: ")
		fmt.Println(data)
		return
	}


	if pedido.EstadoRepartidor == -1{
		m.Message =" el pedido ya fue entreagado "
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


//Recie un struct de pedidoRepartidor, solo para obtener el id del pedido
//actualiza el estado del repartidor en relacion con el pedido
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
		fmt.Print("pedido recibida: ")
		fmt.Println(data)
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

//funcion que expone los servicios del repartidor
func handle()  {

	router := mux.NewRouter()
	router.HandleFunc("/recibir_pedidio",recibir_pedido).Methods("POST")
	router.HandleFunc("/informar_estado_cliente",informar_estado_cliente).Methods("GET")
	router.HandleFunc("/marcar_pedido",marcar_pedido).Methods("POST")
	http.ListenAndServe(":8082", router)

}

func main (){

	println("escuchando el puerto 8082")
	handle()

}
