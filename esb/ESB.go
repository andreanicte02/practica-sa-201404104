package main

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
	data:= utils.JSONMessageGeneric{}
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	fmt.Println("data recibida: ")
	fmt.Println(data)


	padre, existePadre := utils.GetDataService(servicios,data.Message,"solicitar_pedido")

	var m utils.JSONMessageGeneric

	if existePadre {
		m = utils.JSONMessageGeneric{"Si existe el servicio", 1}
	}else{
		m = utils.JSONMessageGeneric{"no existe el servicio", 0}
	}


	dataRespuesta:= utils.PeticionClienteGeneric(&utils.JSONGenerico{data.Id},padre.Method,padre.Host, padre.Ruta)

	fmt.Println("data recibida:")
	fmt.Println(dataRespuesta)

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(m)
}


func handle()  {

	router := mux.NewRouter()

	router.HandleFunc("/cliente_solicitar_pedido",clienteSolicitarPedido).Methods("POST")
	router.HandleFunc("/registrar_microservicio",registrarMicroServicio).Methods("POST")
	http.ListenAndServe(":8085", router)

}


func main()  {

	handle()

}
