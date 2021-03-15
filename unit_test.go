package main

import (
	"./utils"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)
///// Dado que las pruebas unitarias no deben tener ningún tipo de dependencia,


/*

	preuba que verifica si la suma de dos numeros es correcta

 */

func TestSuma(t *testing.T) {
	fmt.Println("Test funcion suma:")
	valor := utils.Suma(7, 23)
	if valor != 30 {
		t.Error("Se esperaba 30 y se obtuvo", valor)
	}
}

/*

	preuba que verifica que el procedimiento de buscar servicios sea correcto

*/

func TestGetDataService(t *testing.T) {
	fmt.Println("Test funcion GetDataService:")
	servicios := []utils.ServicioData{}
	servicios = append(servicios, utils.ServicioData{"8082", "informar_estado_cliente","/informar_estado_cliente","repartidor","GET"})
	servicios = append(servicios, utils.ServicioData{"8081", "estado_pedido","/estado_pedido","restaurante","GET"})

	_,existe := utils.GetDataService(servicios,"repartidor","informar_estado_cliente")

	if !existe {
		t.Error("se esperaba un valor verdero y se obtuvo", false)
	}

}


/*

	preuba que verifica si la funcion para decodifcar funciona de manera cocrrecta

*/

func TestDecodificador(t *testing.T) {
	fmt.Println("Test funcion GetDataService:")
	body := ioutil.NopCloser(strings.NewReader("\"{\"Message\":\"Hola mundo\",\"Id\":1}\""))
	message:= utils.JSONMessageGeneric{"test",1}
	valor:= utils.Decodificador(body,&message)

	if(valor.Message!="test"){
		t.Error("Se esperaba la palabra test y se obutvo", valor.Message)
	}
}