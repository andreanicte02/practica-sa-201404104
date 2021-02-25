package main

import (
	"../models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

//se usan hash para llevar el control de clientes, pedido, menu en memoria
var hashCliente = make(map[int]*models.Cliente)
var hashPedido = make(map[int]*models.Pedido)
var hashMenu = make(map[int]*models.Menu)
//contador globla y unico que va llevar el id de los pedidos
var idPedido = 0

/**
w -> indica el contenido de la respuesta
r -> indica el contenido de la solicitud
*/

//funcion que recibe el pedido y lo guarda en memoria, recibe una estructura de pedido en el body
func recibir_pedidio(w http.ResponseWriter, r *http.Request)  {

	data:= models.Pedido{}
	m:= models.JSONMessageGeneric{}
	w.Header().Set("Content-Type","application/json")

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_,existeCliente := hashCliente[data.IdCliente]
	_,existeMenu := hashMenu[data.IdMenu]

	if !existeCliente || !existeMenu{

		mensaje_error,_ := json.Marshal(models.JSONMessageGeneric{"El menu o el cliente no existen",-1})
		http.Error(w, string(mensaje_error), http.StatusBadRequest)
		return
	}


	//se retorna un json con el id del pedido
	hashPedido[idPedido]=&data
	m.Message =  "pedidio realizado"
	m.Id = idPedido
	json.NewEncoder(w).Encode(m)
	fmt.Print("data recibida: ")
	fmt.Println(data)
	idPedido ++


}

////funcion que recibe el id del pedido y devuelve el estado del pedido
func etado_pedido(w http.ResponseWriter, r *http.Request)  {

	data:= models.JSONGenerico{}
	m:= models.JSONMessageGeneric{}
	w.Header().Set("Content-Type","application/json")
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pedido,existePedido := hashPedido[data.Id]

	if !existePedido{

		mensaje_error,_ := json.Marshal(models.JSONMessageGeneric{"Ese pedido no existe",-1})
		http.Error(w, string(mensaje_error), http.StatusBadRequest)
		return

	}

	var strStatus string

	if strStatus = "pendiente"; pedido.IdEstado == 1{

		strStatus = "completo"
	}


	m.Message = hashCliente[pedido.IdCliente].Nombre + "respuesta estado de pedidio " + strStatus
	m.Id = pedido.IdEstado
	json.NewEncoder(w).Encode(m)

	fmt.Print("data recibida: ")
	fmt.Println(data)

}

////funcion que recibe el id del pedido y indica en memoria y al repartidor que el pedido ya se puede recoger
func avisar_pedido_listo(w http.ResponseWriter, r *http.Request)  {

	data:= models.JSONGenerico{}
	m:= models.JSONMessageGeneric{}
	w.Header().Set("Content-Type","application/json")
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pedido,existePedido := hashPedido[data.Id]

	if !existePedido{

		mensaje_error,_ := json.Marshal(models.JSONMessageGeneric{"Ese pedido no existe",-1})
		http.Error(w, string(mensaje_error), http.StatusBadRequest)
		return

	}

	simulacionEntregaPedidoAlRepartidor(data.Id)
	m.Message = hashCliente[pedido.IdCliente].Nombre + " pedido mas que listo "
	m.Id = 1
	pedido.IdEstado = 1
	json.NewEncoder(w).Encode(m)
	fmt.Print("data recibida: ")
	fmt.Println(data)


}
//fncion que simula la comunicacion entre el servicio del repartidor y el restaruante
func simulacionEntregaPedidoAlRepartidor(idPedido int){


	pedido,existePedido := hashPedido[idPedido]

	if !existePedido || pedido.IdEstado==1{
		return
	}

	data,_:= json.Marshal(models.PedidoRepartidor{pedido.IdMenu,pedido.IdCliente, pedido.IdEstado, hashMenu[pedido.IdMenu].Descripcion,idPedido,0})
	req,err := http.NewRequest("POST", "http://localhost:8082/recibir_pedidio", bytes.NewBuffer(data))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string([]byte(body)))
	pedido.IdEstado = 1



}


//funcion que expone los servicios del restaurante
func handle()  {

	router := mux.NewRouter()
	router.HandleFunc("/recibir_pedido",recibir_pedidio).Methods("POST")
	router.HandleFunc("/estado_pedido",etado_pedido).Methods("GET")
	router.HandleFunc("/avisar_pedido_listo",avisar_pedido_listo).Methods("POST")
	http.ListenAndServe(":8081", router)

}





func main()  {

	hashMenu[0]=&models.Menu{0,"menu1"}
	hashMenu[1]=&models.Menu{1,"menu1"}
	hashMenu[2]=&models.Menu{2,"menu1"}

	hashCliente[0]=&models.Cliente{0,"cliente1"}
	hashCliente[1]=&models.Cliente{1,"cliente2"}
	hashCliente[2]=&models.Cliente{2,"cliente2"}

	println("escuchando el puerto 8081")
	handle()



}