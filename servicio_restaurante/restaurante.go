package main

import (
	"../estructura"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

var hashCliente = make(map[int]estructura.Cliente)
var hashPedido = make(map[int]estructura.Pedido)
var hashMenu = make(map[int]estructura.Menu)
var idPedido = 0


func recibir_pedidio(w http.ResponseWriter, r *http.Request)  {

	data:= estructura.Pedido{}
	m:= estructura.JSONMessageGeneric{}
	w.Header().Set("Content-Type","application/json")

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_,existeCliente := hashCliente[data.IdCliente]
	_,existeMenu := hashMenu[data.IdMenu]

	if !existeCliente || !existeMenu{

		mensaje_error,_ := json.Marshal(estructura.JSONMessageGeneric{"El menu o el cliente no existen",-1})
		http.Error(w, string(mensaje_error), http.StatusBadRequest)
		return
	}


	hashPedido[idPedido]=data
//simulacionEntregaPedidoAlRepartidor(idPedido)
	m.Message =  "pedidio realizado"
	m.Id = idPedido
	json.NewEncoder(w).Encode(m)
	fmt.Print("data recibida: ")
	fmt.Println(data)
	idPedido ++


}
//recibimos el id del cliente
func etado_pedido(w http.ResponseWriter, r *http.Request)  {

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

	var strStatus string

	if strStatus = "pendiente"; pedido.IdEstado == 1{

		strStatus = "completo"
	}


	m.Message = hashCliente[pedido.IdCliente].Nombre + "respuesta estado de pedidio" + strStatus
	m.Id = pedido.IdEstado
	json.NewEncoder(w).Encode(m)

	fmt.Print("data recibida: ")
	fmt.Println(data)

}

func avisar_pedido_listo(w http.ResponseWriter, r *http.Request)  {

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

	simulacionEntregaPedidoAlRepartidor(data.Id)
	m.Message = hashCliente[pedido.IdCliente].Nombre + " pedido mas que listo "
	m.Id = 1
	pedido.IdEstado = 1
	json.NewEncoder(w).Encode(m)
	fmt.Print("data recibida: ")
	fmt.Println(data)


}

func simulacionEntregaPedidoAlRepartidor(idPedido int){


	pedido,existePedido := hashPedido[idPedido]

	if !existePedido || pedido.IdEstado==1{
		return
	}



	data,_:= json.Marshal(estructura.PedidoRepartidor{pedido.IdMenu,pedido.IdCliente, pedido.IdEstado, hashMenu[pedido.IdMenu].Descripcion,idPedido})
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

func handle()  {

	router := mux.NewRouter()
	router.HandleFunc("/crear_pedido",recibir_pedidio).Methods("POST")
	router.HandleFunc("/estado_pedido",etado_pedido).Methods("GET")
	router.HandleFunc("/avisar_pedido_listo",avisar_pedido_listo).Methods("POST")
	http.ListenAndServe(":8081", router)

}





func main()  {

	hashMenu[0]=estructura.Menu{0,"menu1"}
	hashMenu[1]=estructura.Menu{1,"menu1"}
	hashMenu[2]=estructura.Menu{2,"menu1"}

	hashCliente[0]=estructura.Cliente{0,"cliente1"}
	hashCliente[1]=estructura.Cliente{1,"cliente2"}
	hashCliente[2]=estructura.Cliente{2,"cliente2"}


	handle()



}
