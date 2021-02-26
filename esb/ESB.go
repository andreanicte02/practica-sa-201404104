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
	w.Header().Set("Content-Type","application/json")
	m= utils.JSONMessageGeneric{Message: "Servicio Registrado", Id: 1}

	json.NewEncoder(w).Encode(m)
	fmt.Println("Servicio registrado: " + dataServicio.Nombre + "microservicio: " + dataServicio.Padre)
	fmt.Println("data")
	
}




func handle()  {

	router := mux.NewRouter()

	router.HandleFunc("/registrar_microservicio",registrarMicroServicio).Methods("POST")
	http.ListenAndServe(":8085", router)

}


func main()  {

	handle()

}
