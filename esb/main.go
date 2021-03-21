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
var servicios = []ServicioData{}
var m JSONMessageGeneric


func registrarMicroServicio(w http.ResponseWriter, r *http.Request)  {

	var dataServicio = ServicioData{}
	err := json.NewDecoder(r.Body).Decode(&dataServicio)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	servicios = append(servicios, dataServicio)
	m= JSONMessageGeneric{Message: "Servicio Registrado", Id: 1}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(m)
	fmt.Println("Servicio registrado: " + dataServicio.Nombre +  " microservicio: " + dataServicio.Padre)
	fmt.Println("data")
	
}

//end point 1
func clienteSolicitarPedido(w http.ResponseWriter, r *http.Request)  {


	//recibimos la informacion y el padre del servicio en este cado es id-padre
	data:= JSONGenerico{}
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	fmt.Println("data recibida: de solicitar_pedido en cliente")
	fmt.Println(data)


	padre, existePadre := GetDataService(servicios,"cliente","solicitar_pedido")
	if !existePadre{
		fmt.Println("no existe servicio")
		return
	}

	dataRespuesta:= PeticionJSONGeneric(&data,padre.Method,padre.Host, padre.Ruta)

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(dataRespuesta)
}

//endpoint 2

func restauranteRecibirPedido(w http.ResponseWriter, r *http.Request)  {


	//recibimos la informacion y el padre del servicio en este cado es id-padre
	data:= Pedido{}
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	fmt.Println("data recibida de el servicio recibir_pedido de restaurante: ")
	fmt.Println(data)


	padre, existePadre := GetDataService(servicios,"restaurante","recibir_pedido")
	if!existePadre{
		fmt.Println("no existe servicio")
		return
	}


	dataRespuesta:= PeticionRestaurante(&data,padre.Method,padre.Host, padre.Ruta)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(dataRespuesta)

}


//endpoint 3
func clienteEstadoRestaurante(w http.ResponseWriter, r *http.Request)  {


	//recibimos la informacion y el padre del servicio en este cado es id-padre
	data:= JSONGenerico{}
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	fmt.Println("data recibida: de  get_estado_restaurante cliente")
	fmt.Println(data)


	padre, existePadre :=GetDataService(servicios,"cliente","get_estado_restaurante")
	if!existePadre{
		fmt.Println("no existe servicio")
		return
	}


	dataRespuesta:= PeticionJSONGeneric(&data,padre.Method,padre.Host, padre.Ruta)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(dataRespuesta)

}


//endpoint 4
func restauranteEstadoPedido(w http.ResponseWriter, r *http.Request)  {


	//recibimos la informacion y el padre del servicio en este cado es id-padre
	data:=JSONGenerico{}
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	fmt.Println("data recibida: de estado_pedido de restaurante")
	fmt.Println(data)


	padre, existePadre := GetDataService(servicios,"restaurante","estado_pedido")
	if!existePadre{
		fmt.Println("no existe servicio")
		return
	}


	dataRespuesta:= PeticionJSONGeneric(&data,padre.Method,padre.Host, padre.Ruta)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(dataRespuesta)

}

//endpoint 5
func clienteEstadoRepartidor(w http.ResponseWriter, r *http.Request)  {


	//recibimos la informacion y el padre del servicio en este cado es id-padre
	data:= JSONGenerico{}
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	fmt.Println("data recibida: get_estado_repartidor en cliente ")
	fmt.Println(data)


	padre, existePadre := GetDataService(servicios,"cliente","get_estado_repartidor")
	if!existePadre{
		fmt.Println("no existe servicio")
		return
	}


	dataRespuesta:= PeticionJSONGeneric(&data,padre.Method,padre.Host, padre.Ruta)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(dataRespuesta)

}

//endpoint 6
func repartidorEstado(w http.ResponseWriter, r *http.Request)  {


	//recibimos la informacion y el padre del servicio en este cado es id-padre
	data:= JSONGenerico{}
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	fmt.Println("data recibida: de repartidor de informar_estado_cliente ")
	fmt.Println(data)


	padre, existePadre := GetDataService(servicios,"repartidor","informar_estado_cliente")
	if!existePadre{
		fmt.Println("no existe servicio")
		return
	}


	dataRespuesta:= PeticionJSONGeneric(&data,padre.Method,padre.Host, padre.Ruta)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(dataRespuesta)

}

//endpoint 7
func restaurantePedidoListo(w http.ResponseWriter, r *http.Request)  {


	//recibimos la informacion y el padre del servicio en este cado es id-padre
	data:= JSONGenerico{}
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	fmt.Println("data recibida: de restaurante de avisar_pedido_listo ")
	fmt.Println(data)


	padre, existePadre := GetDataService(servicios,"restaurante","avisar_pedido_listo")
	if!existePadre{
		fmt.Println("no existe servicio")
		return
	}


	dataRespuesta:= PeticionJSONGeneric(&data,padre.Method,padre.Host, padre.Ruta)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(dataRespuesta)

}

//endpoint 8
func repartidorRecibirPedido(w http.ResponseWriter, r *http.Request)  {


	//recibimos la informacion y el padre del servicio en este cado es id-padre
	data:= PedidoRepartidor{}
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	fmt.Println("data recibida: de repartidor de recibir_pedidio")
	fmt.Println(data)


	padre, existePadre := GetDataService(servicios,"repartidor","recibir_pedidio")
	if!existePadre{
		fmt.Println("no existe servicio")
		return
	}


	dataRespuesta:= PeticionRepartodpr(&data,padre.Method,padre.Host, padre.Ruta)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(dataRespuesta)

}
//endpoint9
func repartidorMarcarPedido(w http.ResponseWriter, r *http.Request)  {


	//recibimos la informacion y el padre del servicio en este cado es id-padre
	data:= JSONGenerico{}
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	fmt.Println("data recibida: de repartidor de marcar_pedido ")
	fmt.Println(data)


	padre, existePadre := GetDataService(servicios,"repartidor","marcar_pedido")
	if!existePadre{
		fmt.Println("no existe servicio")
		return
	}


	dataRespuesta:= PeticionJSONGeneric(&data,padre.Method,padre.Host, padre.Ruta)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(dataRespuesta)

}


func handle()  {

	router := mux.NewRouter()


	router.HandleFunc("/repartidor_marcar_pedido",repartidorMarcarPedido).Methods("POST")
	router.HandleFunc("/repartidor_recibir_pedidio",repartidorRecibirPedido).Methods("POST")
	router.HandleFunc("/restaurante_pedido_listo",restaurantePedidoListo).Methods("POST")
	router.HandleFunc("/repartidor_estado",repartidorEstado).Methods("GET")
	router.HandleFunc("/cliente_estado_repartidor",clienteEstadoRepartidor).Methods("GET")
	router.HandleFunc("/restaurante_estado_restaurante",restauranteEstadoPedido).Methods("GET")
	router.HandleFunc("/cliente_estado_restaurante",clienteEstadoRestaurante).Methods("GET")
	router.HandleFunc("/cliente_solicitar_pedido",clienteSolicitarPedido).Methods("POST")
	router.HandleFunc("/restaurante_recibir_pedido",restauranteRecibirPedido).Methods("POST")
	router.HandleFunc("/registrar_microservicio",registrarMicroServicio).Methods("POST")
	http.ListenAndServe(":8085", router)

}


func main()  {

	fmt.Println("escuhando el puerto 8085")
	handle()

}
