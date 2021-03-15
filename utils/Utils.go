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

	Message string `json:"message"`
	Id      int `json:"id"`

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

//decodificador
func Decodificador(Body io.ReadCloser, Data *JSONMessageGeneric) JSONMessageGeneric  {
	decoder:= json.NewDecoder(Body)
	decoder.Decode(Data)
	return *Data
}


//funcion para registrar servicios
func RegistrarServicio(servicio *ServicioData, method string, host string, nameSerivce string){

	dataRequest,_:= json.Marshal(servicio)
	req,err := http.NewRequest(method, "http://localhost:"+host+ nameSerivce, bytes.NewBuffer(dataRequest))
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


//buscar en array
func GetDataService(array []ServicioData, padre string, nombreServicio string) (ServicioData, bool){

	for i:= 0; i< len(array); i++{

		if array[i].Padre==padre && array[i].Nombre == nombreServicio {

			return array[i],true
		}

	}
	return ServicioData{"","","","",""},false

}

func PeticionJSONGeneric(servicio *JSONGenerico, method string, host string, rutaServicio string) JSONMessageGeneric {

	dataRequest, _ := json.Marshal(servicio)
	req, err := http.NewRequest(method, "http://localhost:"+host+rutaServicio, bytes.NewBuffer(dataRequest))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	defer resp.Body.Close()
	var data = Decodificador(resp.Body, &JSONMessageGeneric{"", 0})
	return data

}

func PeticionRestaurante(servicio *Pedido, method string, host string, rutaServicio string) JSONMessageGeneric {

	dataRequest, _ := json.Marshal(servicio)
	req, err := http.NewRequest(method, "http://localhost:"+host+rutaServicio, bytes.NewBuffer(dataRequest))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	defer resp.Body.Close()
	var data = Decodificador(resp.Body, &JSONMessageGeneric{"", 0})
	return data

}

func PeticionRepartodpr(servicio *PedidoRepartidor, method string, host string, rutaServicio string) JSONMessageGeneric {

	dataRequest, _ := json.Marshal(servicio)
	req, err := http.NewRequest(method, "http://localhost:"+host+rutaServicio, bytes.NewBuffer(dataRequest))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	defer resp.Body.Close()
	var data = Decodificador(resp.Body, &JSONMessageGeneric{"", 0})

	return data

}

func Suma(numero1, numero2 int) (resultado int) {
	resultado = numero1 + numero2
	return
}