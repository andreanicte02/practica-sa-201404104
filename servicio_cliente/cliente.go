
package main

import (
	"../estructura"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
)


var hashPedido = make(map[int]int)
var codigo int = -1

func solicitar_pedido(pedido *estructura.Pedido){
	
	dataRequest,_:= json.Marshal(pedido)
	req,err := http.NewRequest("POST", "http://localhost:8081/recibir_pedido", bytes.NewBuffer(dataRequest))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	defer resp.Body.Close()

	var data estructura.JSONMessageGeneric
	decoder:= json.NewDecoder(resp.Body)
	decoder.Decode(&data)

	hashPedido[data.Id]=data.Id
	fmt.Println("...respuesta")
	fmt.Println(data)
	fmt.Println("id del pedido:")
	fmt.Println(data.Id)


}


func solicitar_estado_restaurante(jsonGeneric *estructura.JSONGenerico){

	dataRequest,_:= json.Marshal(jsonGeneric)
	req,err := http.NewRequest("GET", "http://localhost:8081/estado_pedido", bytes.NewBuffer(dataRequest))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	defer resp.Body.Close()

	var data estructura.JSONMessageGeneric
	decoder:= json.NewDecoder(resp.Body)
	decoder.Decode(&data)

	hashPedido[data.Id]=data.Id
	fmt.Println("...respuesta")
	fmt.Println(data)


}


func estado_repartidor(jsonGeneric *estructura.JSONGenerico){

	dataRequest,_:= json.Marshal(jsonGeneric)
	req,err := http.NewRequest("GET", "http://localhost:8082/informar_estado_cliente", bytes.NewBuffer(dataRequest))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	defer resp.Body.Close()

	var data estructura.JSONMessageGeneric
	decoder:= json.NewDecoder(resp.Body)
	decoder.Decode(&data)

	hashPedido[data.Id]=data.Id
	fmt.Println("...respuesta")
	fmt.Println(data)


}


func menuRandom()  int {

	return rand.Intn(2 - 0) + 0
}


func handle() {

	router := mux.NewRouter()
	http.ListenAndServe(":8080", router)
}


func menu(){

	option:=-1
	idPedido:= -1
	for true {


		fmt.Println("Menu acciones:")
		fmt.Println("1. Solicitar pedido")
		fmt.Println("2. Solicitar estado pedio al restaurante")
		fmt.Println("3. Solicitar estado del repartidor")
		fmt.Scanf("%d", &option)

		switch option {

		case 1:

			println(".....solicitando pedido")
			solicitar_pedido(&estructura.Pedido{menuRandom(),codigo,0})

			break

		case 2:
			println("Ingrese.....codigo_pedido:")
			fmt.Scanf("%d", &idPedido)

			println(".....solicitando estado restaurante")
			solicitar_estado_restaurante(&estructura.JSONGenerico{idPedido})

			break

		case 3:

			println("Ingrese.....codigo_pedido:")
			fmt.Scanf("%d", &idPedido)

			println(".....solicitando estado repartidor")
			estado_repartidor(&estructura.JSONGenerico{idPedido})
			break


		}


	}
}


func main()  {

	codigo=-1


	for true{

		fmt.Println("Ingrse el nombre del cliente:")
		fmt.Scanf("%d", &codigo)
		if codigo > 2{
			fmt.Println("no existe el cliente en sistema")
			continue
		}
		menu()

	}


}
