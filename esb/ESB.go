package esb

import (
	"../utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var servicios = []utils.ServicioData{}
var m utils.JSONMessageGeneric


func registrarMicroServicio(w http.ResponseWriter, r *http.Request)  {

	var dataServicio = utils.ServicioData{}
	err := json.NewDecoder(r.Body).Decode(&dataServicio)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	servicios = append(servicios, dataServicio)
	m= utils.JSONMessageGeneric{Message: "Servicio Registrado", Id: 1}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(m)
	fmt.Println("Servicio registrado: " + dataServicio.Nombre +  " microservicio: " + dataServicio.Padre)
	fmt.Println("data")
	
}

//end point 1
func clienteSolicitarPedido(w http.ResponseWriter, r *http.Request)  {


	//recibimos la informacion y el padre del servicio en este cado es id-padre
	data:= utils.JSONGenerico{}
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	fmt.Println("data recibida: de solicitar_pedido en cliente")
	fmt.Println(data)


	padre, existePadre := utils.GetDataService(servicios,"cliente","solicitar_pedido")
	if !existePadre{
		fmt.Println("no existe servicio")
		return
	}

	dataRespuesta:= utils.PeticionJSONGeneric(&data,padre.Method,padre.Host, padre.Ruta)

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(dataRespuesta)
}

//endpoint 2

func restauranteRecibirPedido(w http.ResponseWriter, r *http.Request)  {


	//recibimos la informacion y el padre del servicio en este cado es id-padre
	data:= utils.Pedido{}
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	fmt.Println("data recibida de el servicio recibir_pedido de restaurante: ")
	fmt.Println(data)


	padre, existePadre := utils.GetDataService(servicios,"restaurante","recibir_pedido")
	if!existePadre{
		fmt.Println("no existe servicio")
		return
	}


	dataRespuesta:= utils.PeticionRestaurante(&data,padre.Method,padre.Host, padre.Ruta)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(dataRespuesta)

}


//endpoint 3
func clienteEstadoRestaurante(w http.ResponseWriter, r *http.Request)  {


	//recibimos la informacion y el padre del servicio en este cado es id-padre
	data:= utils.JSONGenerico{}
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	fmt.Println("data recibida: de  get_estado_restaurante cliente")
	fmt.Println(data)


	padre, existePadre := utils.GetDataService(servicios,"cliente","get_estado_restaurante")
	if!existePadre{
		fmt.Println("no existe servicio")
		return
	}


	dataRespuesta:= utils.PeticionJSONGeneric(&data,padre.Method,padre.Host, padre.Ruta)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(dataRespuesta)

}


//endpoint 4
func restauranteEstadoPedido(w http.ResponseWriter, r *http.Request)  {


	//recibimos la informacion y el padre del servicio en este cado es id-padre
	data:= utils.JSONGenerico{}
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	fmt.Println("data recibida: de estado_pedido de restaurante")
	fmt.Println(data)


	padre, existePadre := utils.GetDataService(servicios,"restaurante","estado_pedido")
	if!existePadre{
		fmt.Println("no existe servicio")
		return
	}


	dataRespuesta:= utils.PeticionJSONGeneric(&data,padre.Method,padre.Host, padre.Ruta)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(dataRespuesta)

}

//endpoint 5
func clienteEstadoRepartidor(w http.ResponseWriter, r *http.Request)  {


	//recibimos la informacion y el padre del servicio en este cado es id-padre
	data:= utils.JSONGenerico{}
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	fmt.Println("data recibida: get_estado_repartidor en cliente ")
	fmt.Println(data)


	padre, existePadre := utils.GetDataService(servicios,"cliente","get_estado_repartidor")
	if!existePadre{
		fmt.Println("no existe servicio")
		return
	}


	dataRespuesta:= utils.PeticionJSONGeneric(&data,padre.Method,padre.Host, padre.Ruta)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(dataRespuesta)

}

//endpoint 6
func repartidorEstado(w http.ResponseWriter, r *http.Request)  {


	//recibimos la informacion y el padre del servicio en este cado es id-padre
	data:= utils.JSONGenerico{}
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	fmt.Println("data recibida: de repartidor de informar_estado_cliente ")
	fmt.Println(data)


	padre, existePadre := utils.GetDataService(servicios,"repartidor","informar_estado_cliente")
	if!existePadre{
		fmt.Println("no existe servicio")
		return
	}


	dataRespuesta:= utils.PeticionJSONGeneric(&data,padre.Method,padre.Host, padre.Ruta)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(dataRespuesta)

}

//endpoint 7
func restaurantePedidoListo(w http.ResponseWriter, r *http.Request)  {


	//recibimos la informacion y el padre del servicio en este cado es id-padre
	data:= utils.JSONGenerico{}
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	fmt.Println("data recibida: de restaurante de avisar_pedido_listo ")
	fmt.Println(data)


	padre, existePadre := utils.GetDataService(servicios,"restaurante","avisar_pedido_listo")
	if!existePadre{
		fmt.Println("no existe servicio")
		return
	}


	dataRespuesta:= utils.PeticionJSONGeneric(&data,padre.Method,padre.Host, padre.Ruta)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(dataRespuesta)

}

//endpoint 8
func repartidorRecibirPedido(w http.ResponseWriter, r *http.Request)  {


	//recibimos la informacion y el padre del servicio en este cado es id-padre
	data:= utils.PedidoRepartidor{}
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	fmt.Println("data recibida: de repartidor de recibir_pedidio")
	fmt.Println(data)


	padre, existePadre := utils.GetDataService(servicios,"repartidor","recibir_pedidio")
	if!existePadre{
		fmt.Println("no existe servicio")
		return
	}


	dataRespuesta:= utils.PeticionRepartodpr(&data,padre.Method,padre.Host, padre.Ruta)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(dataRespuesta)

}
//endpoint9
func repartidorMarcarPedido(w http.ResponseWriter, r *http.Request)  {


	//recibimos la informacion y el padre del servicio en este cado es id-padre
	data:= utils.JSONGenerico{}
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	fmt.Println("data recibida: de repartidor de marcar_pedido ")
	fmt.Println(data)


	padre, existePadre := utils.GetDataService(servicios,"repartidor","marcar_pedido")
	if!existePadre{
		fmt.Println("no existe servicio")
		return
	}


	dataRespuesta:= utils.PeticionJSONGeneric(&data,padre.Method,padre.Host, padre.Ruta)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(dataRespuesta)

}


func Handle()  {

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


