package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)




//struct de un json generico
type JSONGenerico struct {

	Id int `json:"id"`

}
//struct de un mensaje generico
type JSONMessageGeneric struct {

	Message string
	Id      int

}

//json para enviar la data y registrar los servicios en el ESB
type JSONMessageServices struct {

	Name string
	Ruta string
	HOST  int

}


//struct que nos va ayudar a simular los menus en memoria
type Menu struct {
	Id int
	Descripcion string
}

//struct que nos va ayudar a simular los clientes en memoria
type Cliente struct {
	Id int
	Nombre string
}

//struct que va ayudar a almacenar la informacion de pedidos en memoria del restaurante
type Pedido struct {

	IdMenu int `json:"idMenu"`
	IdCliente int `json:"idCliente"`
	IdEstado int `json:"IdEstado"` //0 pendiente 1 completado


}


//struct que va ayudar a almacenar la informacion de pedidos en memoria del repartidor
type PedidoRepartidor struct {

	IdMenu int `json:"idMenu"`
	IdCliente int `json:"idCliente"`
	IdEstado int `json:"IdEstado"` //0 pendiente 1 completado
	DescripcionMenu string `json:"DescripcionMenu"`
	IdPedido int  `json:"idPedido"`
	EstadoRepartidor int `json:"estadoRepartidor ya tomo el pedido o no"`


}

func LogSalida(data JSONGenerico, m JSONMessageGeneric)  {
	fmt.Println("data recibida: ")
	fmt.Println(data)
	fmt.Println("respuesta de la solicitud de pedido:")
	fmt.Println(m)
}


type ServicioData struct {

	Host string
	Nombre string
	Ruta string
	Padre string
	Method string


}

func Decodificador(body io.ReadCloser, data *JSONMessageGeneric) JSONMessageGeneric  {
	decoder:= json.NewDecoder(body)
	decoder.Decode(data)
	return *data
}


func RegistrarServicio(servicio *ServicioData, method string,){

	dataRequest,_:= json.Marshal(servicio)
	req,err := http.NewRequest(method, "http://localhost:8085/registrar_microservicio", bytes.NewBuffer(dataRequest))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	defer resp.Body.Close()
	var data = Decodificador(resp.Body,&JSONMessageGeneric{"",0})
	fmt.Println("info. recibida")
	fmt.Println(data)

}


