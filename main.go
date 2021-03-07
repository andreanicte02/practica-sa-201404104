package main

import (
	"./esb"
	"./serivicio-cliente"
	"./servicio-repartidor"
	"./servicio-restaruante"
	"fmt"
)

func main()  {


	go func() {
		fmt.Println("escuhando el puerto 8085")
		esb.Handle()
	}()

	go func() {
		fmt.Println("escuhando el puerto 8085")
		servicio_cliente.Handle()
	}()


	go func() {

		println("escuchando el puerto 8082")
		servicio_repartidor.Handle()
	}()


	println("escuchando el puerto 8081")
	servicio_restaruante.Handle()

}