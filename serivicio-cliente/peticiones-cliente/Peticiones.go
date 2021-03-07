package peticiones_cliente

import (
	"../../utils"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)
var HashPedido = make(map[int]int)
var Codigo  = -1
var IdPedido = -1
/*
	req,err :=  http.NewRequest(1, 2, 3)
		primer argumento, se indica el metodo a utilizar POST | GET | PUT etc
		segundo argumetno, se indica el url de la api a consumir
		tercer agumento, se indica el contenido del pody de la peticion
*/

/*
	agregar informacion al header de una peticion
	req.Header.Add(1, 2)
		primer argumento indica el tipo que se va agregar
		segundo argument indica el valor
*/


//funcion para solicitar pedido
func PeticionSolicitarPedido(pedido *utils.Pedido) utils.JSONMessageGeneric{

	dataRequest,_:= json.Marshal(pedido)
	req,err := http.NewRequest("POST", "http://localhost:8085/restaurante_recibir_pedido", bytes.NewBuffer(dataRequest))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	defer resp.Body.Close()

	var data = utils.Decodificador(resp.Body,&utils.JSONMessageGeneric{"",0})

	HashPedido[data.Id]=data.Id
	IdPedido = data.Id
	return data


}

//funcion para solicitar pedido al restaurante
func PeticionSolicitarEstadoRestaurante(jsonGeneric *utils.JSONGenerico) utils.JSONMessageGeneric{

	dataRequest,_:= json.Marshal(jsonGeneric)
	req,err := http.NewRequest("GET", "http://localhost:8085/restaurante_estado_restaurante", bytes.NewBuffer(dataRequest))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	defer resp.Body.Close()

	var data = utils.Decodificador(resp.Body,&utils.JSONMessageGeneric{"",0})

	HashPedido[data.Id]=data.Id


	return data



}

//funcion para solicitar pedido al repartidor
func PeticionEstadoRepartidor(jsonGeneric *utils.JSONGenerico) utils.JSONMessageGeneric{

	dataRequest,_:= json.Marshal(jsonGeneric)
	req,err := http.NewRequest("GET", "http://localhost:8085/repartidor_estado", bytes.NewBuffer(dataRequest))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	defer resp.Body.Close()


	var data = utils.Decodificador(resp.Body,&utils.JSONMessageGeneric{"",0})

	HashPedido[data.Id]=data.Id
	return data

}

