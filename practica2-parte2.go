package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


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

func rest(body string, url string) {

	req,err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	//el tipo que se envia en el body sera text/xml
	req.Header.Add("Content-Type", "text/xml")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}
	response, _ := ioutil.ReadAll(resp.Body)
	log.Println("....respuesta")
	log.Println(string([]byte(response)))

}

/*
	funcion la que se genra el body que se enviara por la repeticion
	-num1 y num2 seran los valores que se van a operar
	-tipo indicia el tipo de operacion:
		.add
		.divide
		.multiply
		.subtract
*/
func bodySOAPadd(num1 string, num2 string, tipo string) string {

	return "<?xml version=\"1.0\" encoding=\"utf-8\"?>\n<soap12:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap12=\"http://www.w3.org/2003/05/soap-envelope\">\n  <soap12:Body>\n    <"+tipo+" xmlns=\"http://tempuri.org/\">\n      <intA>"+num1+"</intA>\n      <intB>"+num2+"</intB>\n    </"+tipo+">\n  </soap12:Body>\n</soap12:Envelope>"

}

//inicio del flujo de aplicacion
func main()  {
	println("hola mundozz")


	var option int
	var num1 = "dasd"
	var num2 string
	url:= "http://www.dneonline.com/calculator.asmx"

	for true{

		fmt.Println(">>>>>>>>>>>>>>>>>>>>menu<<<<<<<<<<<<<<<<<<<<<<")
		fmt.Println("1. add")
		fmt.Println("2. divide")
		fmt.Println("3. multiply")
		fmt.Println("4. subtract")
		fmt.Println("--elegir la operacion:")
		fmt.Scanf("%d", &option)

		fmt.Println("ingrese num1")
		//se lee el primer numero que se va operar
		fmt.Scanf("%s", &num1)
		//se lee el segundo numero que se va operar
		fmt.Println("ingrese num2")
		fmt.Scanf("%s", &num2)

		switch option {
		case 1:
			rest(bodySOAPadd(num1,num2, "Add"),url)
			break

		case 2:
			rest(bodySOAPadd(num1,num2, "Divide"),url)
			break

		case 3:
			rest(bodySOAPadd(num1,num2, "Multiply"),url)
			break

		case 4:
			rest(bodySOAPadd(num1,num2, "Subtract"),url)
			break
		}


		fmt.Scanf("%d", &option)

	}
}
