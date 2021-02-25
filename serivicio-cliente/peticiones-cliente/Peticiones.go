package peticiones_cliente

import (
	"../../models"
	"bytes"
	"encoding/json"
	"io"
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
func Peticion_solicitar_pedido(pedido *models.Pedido) models.JSONMessageGeneric{

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

	var data = decodificador(resp.Body,&models.JSONMessageGeneric{"",0})

	HashPedido[data.Id]=data.Id
	IdPedido = data.Id
	return data


}

//funcion para solicitar pedido al restaurante
func Peticion_solicitar_estado_restaurante(jsonGeneric *models.JSONGenerico) models.JSONMessageGeneric{

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

	var data = decodificador(resp.Body,&models.JSONMessageGeneric{"",0})

	HashPedido[data.Id]=data.Id


	return data



}

//funcion para solicitar pedido al repartidor
func Peticion_estado_repartidor(jsonGeneric *models.JSONGenerico) models.JSONMessageGeneric{

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


	var data = decodificador(resp.Body,&models.JSONMessageGeneric{"",0})

	HashPedido[data.Id]=data.Id
	return data

}

func decodificador(body io.ReadCloser, data *models.JSONMessageGeneric) models.JSONMessageGeneric  {
	decoder:= json.NewDecoder(body)
	decoder.Decode(data)
	return *data
}