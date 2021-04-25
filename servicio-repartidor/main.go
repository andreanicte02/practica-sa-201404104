package main

import (

	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func Suma(numero1, numero2 int) (resultado int) {
	resultado = numero1 + numero2
	return
}



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



type ServicioData struct {

	Host string
	Nombre string
	Ruta string
	Padre string
	Method string


}

//decodificador
func Decodificador(body io.ReadCloser, data *JSONMessageGeneric) JSONMessageGeneric  {
	decoder:= json.NewDecoder(body)
	decoder.Decode(data)
	return *data
}


//funcion para registrar servicios
func RegistrarServicio(servicio *ServicioData, method string, host string, nameSerivce string){

	dataRequest,_:= json.Marshal(servicio)
	req,err := http.NewRequest(method, "http://esb:"+host+ nameSerivce, bytes.NewBuffer(dataRequest))
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


var estadoRepartidor int =0 //-1 indica que no tiene pedido pendiente, 0 que esta en caminio y 1 que ya se entrego
//se isa una hash en el lado del reapartidor para llevarla en memoria
var hashPedido = make(map[int]*PedidoRepartidor)

/**
w -> indica el contenido de la respuesta
r -> indica el contenido de la solicitud
*/


//servicio que recibe el pedido, indica al reapartidor que tiene pedidos que entregar
//recibe una json con la estrucutra de PedidoRepartidor y actualiza el estado
func recibirPedido(w http.ResponseWriter, r *http.Request)  {

	data:= PedidoRepartidor{}

	m:= JSONMessageGeneric{}
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
func informarEstadoCliente(w http.ResponseWriter, r *http.Request)  {

	data:= JSONGenerico{}

	m:= JSONMessageGeneric{}
	w.Header().Set("Content-Type","application/json")

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pedido,existePedido := hashPedido[data.Id]

	if !existePedido{
		mensaje_error,_ := json.Marshal(JSONMessageGeneric{"El repartidor sigue en espera del pedido",-1})
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
func marcarPedido(w http.ResponseWriter, r *http.Request)  {

	data:= PedidoRepartidor{}
	m:= JSONMessageGeneric{}
	w.Header().Set("Content-Type","application/json")

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pedido,existePedido := hashPedido[data.IdPedido]

	if !existePedido{
		mensaje_error,_ := json.Marshal(JSONMessageGeneric{"Ese pedido no existe",-1})
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
	router.HandleFunc("/recibir_pedidio",recibirPedido).Methods("POST")
	router.HandleFunc("/informar_estado_cliente",informarEstadoCliente).Methods("GET")
	router.HandleFunc("/marcar_pedido",marcarPedido).Methods("POST")
	http.ListenAndServe(":8082", router)

}

func main (){

	RegistrarServicio(&ServicioData{"8082", "recibir_pedidio","/recibir_pedidio","repartidor","POST"}, "POST","8085","/registrar_microservicio")
	RegistrarServicio(&ServicioData{"8082", "informar_estado_cliente","/informar_estado_cliente","repartidor","GET"}, "POST","8085","/registrar_microservicio")
	RegistrarServicio(&ServicioData{"8082", "marcar_pedido","/marcar_pedido","repartidor","POST"}, "POST","8085","/registrar_microservicio")

	println("escuchando el puerto 8082")
	Suma(1,2)
	handle()

}
