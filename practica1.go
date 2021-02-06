package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)


/*
	struct utilizado para manejar la informacion del usuario
*/
type userStruct struct {
	Name string  `json:"name"`
	Gender string `json:"gender"`
	Email string `json:"email"`
	Status string `json:"status"`
}

/*
	struct utilizado para guarda informacion imporante como:
		url de la api
		token de acceso
		id de el usuario

 */

type info struct {
	url string
	tokenB string
	id string
}

////////////////////  seccion de peticiones /////////////////////////////////////////

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


/*
	funcion utilizada para crear un usuario
*/
func createUser(userData []byte, data **info) {

	req,err := http.NewRequest("POST", (*data).url, bytes.NewBuffer(userData))
	req.Header.Set("Authorization", (*data).tokenB)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string([]byte(body)))


}

/*
	funcion utilizada para obtener un usuario
*/
func getUser(data **info) {

	req,err := http.NewRequest("GET", (*data).url+"/"+(*data).id, nil)
	req.Header.Set("Authorization", (*data).tokenB)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string([]byte(body)))

}

/*
	funcion utilizada para actualizar un usuario
*/
func updateUser(userData []byte, data **info) {

	req,err := http.NewRequest("PUT", (*data).url+"/"+(*data).id, bytes.NewBuffer(userData))
	req.Header.Set("Authorization", (*data).tokenB)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string([]byte(body)))

}


/*
	funcion utilizada para eliminar un usuario
*/
func deletUser(data **info) {

	req,err := http.NewRequest("DELETE", (*data).url+"/"+(*data).id, nil)
	req.Header.Set("Authorization", (*data).tokenB)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string([]byte(body)))

}


//Inicio del flujo
func main() {

	//se indica el url del serivico a consumir
	//se configura el token de acessso
	data := info{url: "https://gorest.co.in/public-api/users", tokenB: "Bearer 34607adde1718faea5a4c5f0c73c69c1f864b318679ce0a93dbce6a5e55b907c" , id: "0000"}

	//solamente se incializa el struct, que serivira para guardar la informacion
	user := userStruct {Name: "namexxxx",Gender: "Female", Email: "andrea201404104@email.com", Status: "Active"}

	//inicio del flujo principal.
	menu(&data, &user)
}

func jsonTransform(user **userStruct) []byte {

	jsonUser, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		errors.New("error al convertir a json")
	}

	return jsonUser


}

//funcion que permite codificar informacion json, para mandarlo en el body de la peticion
func menu(data * info, user *userStruct)  {
	scanner := bufio.NewScanner(os.Stdin)
	option := 1
	noValue := 0

	for true {

		fmt.Println(">>>>>>>>>>>>>>menu<<<<<<<<<<<<<<<")
		fmt.Println("1. Crear usuario")
		fmt.Println("2. Obtener usuario")
		fmt.Println("3. Modificar usuario")
		fmt.Println("4. Eliminar usuario")
		fmt.Println("5. salir")
		fmt.Println(">>>>>>>>>>>>>>>><<<<<<<<<<<<<<<<<")

		fmt.Scanf("%d", &option)

		switch option {

		case 1:

			fmt.Println("Ingresar nombre y carnet:")
			scanner.Scan()
			user.Name = scanner.Text()

			fmt.Println("Ingrese email:")
			scanner.Scan()
			user.Email = scanner.Text()

			fmt.Println("Response:......")

			//se obtiene la informacion necesaria, se crea un usuario
			createUser(jsonTransform(&user), &data)

			fmt.Scanf("%d", &noValue)
			break

		case 2:
			fmt.Println("Ingrese el codigo:")

			fmt.Scanln(&data.id)

			fmt.Println("Response:......")

			//se obtiene la informacion necesaria, se obtiene un usuario
			getUser(&data)



			fmt.Scanf("%d", &noValue)

			break
		case 3:

			fmt.Println("Ingrese el nombre a cambiar, correo y estado (Active | Inactive )")
			scanner.Scan()
			user.Name= scanner.Text()

			scanner.Scan()
			user.Email= scanner.Text()

			scanner.Scan()
			user.Status= scanner.Text()

			fmt.Println("Ingrese el codigo:")
			fmt.Scanln(&data.id)


			//se obtiene la informacion necesaria, se actualiza un usuario
			updateUser(jsonTransform(&user),&data)
			fmt.Println("Response:......")
			fmt.Scanf("%d", &noValue)

			break

		case 4:

			fmt.Println("Ingrese el codigo:")

			fmt.Scanln(&data.id)

			fmt.Println("Response:......")
			//se obtiene la informacion necesaria, se elimina un usuario
			deletUser(&data)

			fmt.Scanf("%d", &noValue)


			break
		case 5:
			return




		}


	}

}



