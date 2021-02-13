package main

import (
	"bufio"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	-struct utilizado para almacenar los datos del usuario
		--nombre
		--carnet

 */
type user struct {

	Name string
	Carne int

}

/*
	-struct urilido para guarda la informacion del header

*/

type header struct {


	Alg string `json:"alg"`
	Typ string `json:"type"`
}

/*

	-struct utilizado para guardar los daots del usuario y su respectivo secrect

 */

type usersStruct struct {

	user *user
	secret string

}


/*
	-funcion para codificar string en base 64
*/
func base64Encode(src string) string {
	return strings.
		TrimRight(base64.URLEncoding.
			EncodeToString([]byte(src)), "=")
}

/*
	-funcion para decodficar string en base 64
*/
func base64Decode(src string) string {
	if l:= len(src) % 4; l > 0 {
		src += strings.Repeat("=", 4-l)
	}
	decoded, _ := base64.URLEncoding.DecodeString(src)

	return string(decoded)
}

/*
	-funcion para generar el secreto
*/
func generateScret() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}




func main() {

	menu()


}

/*
	funcion que permite generar la parte de  SIGNATURE del jwt

 */
func generateSignature(secret string, strFinal string) string {

	key:= []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(strFinal))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))

	
}

/*
	funcion que permite generar to-do el string del token jwt

*/
func generateJWT(user * user, secret string) string  {

	headerData := header {"HS256", "JWT"}
	jsonHeader, _ := json.Marshal(headerData)

	encodedHeader:= base64Encode(string(jsonHeader))

	jsonPayload,_ := json.Marshal(user)
	encodedPyload := base64Encode(string(jsonPayload))

	strFinal := encodedHeader+"."+encodedPyload

	return strFinal+"."+generateSignature(secret,strFinal)



}

/*
	130-esta funcion decodifica el la parte del token pyload
	135-verifica si existe en memoria el usuario con ese numero de carnet
	140-si si existe procede a generar to-do el token nuevamente, con el pyload codificado
	142-si es igual el token es aprovado
	147-si no es denegado
*/
func confirmToken(jwt string, userStruct *usersStruct) {
	arrayJWT := strings.Split(jwt, ".")

	if len(arrayJWT) != 3{
		fmt.Println("token denegado")
		return
	}

	//la data pyload
	jsonPyload := base64Decode(arrayJWT[1])
	var user user
	json.Unmarshal([]byte(jsonPyload), &user)

	if user.Carne != userStruct.user.Carne{
		fmt.Println("el usuario con ese carnet no tiene permiso")
		return
	}

	strJWTGeneradoDeNuevo := generateJWT(&user, userStruct.secret)

	if jwt == strJWTGeneradoDeNuevo {
		fmt.Println("Acceso conseguido")
		fmt.Println("jwt generado primero"+ jwt)
		fmt.Println("jwt generado para comparar"+ strJWTGeneradoDeNuevo)

	}else{
		fmt.Println("Acceso denegado")
	}

}

/*

	-funcion que inicia el flujo de la aplicacion

*/

func menu(){


	for true {

		var user user
		//se genra secret
		secret := generateScret()
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println(">>>>>>>>>>>>>>>>>>>>menu<<<<<<<<<<<<<<<<<<<<<<")
		fmt.Println("Ingresar nombre:")

		//se empiezan a capturar los datos de nombre y usuario
		scanner.Scan()
		user.Name = scanner.Text()

		fmt.Println("Carnet:")

		scanner.Scan()
		user.Carne, _ = strconv.Atoi(scanner.Text())

		//se gemera el string de jwt
		jwt :=generateJWT(&user, secret)
		fmt.Println("Token generado: "+ jwt)

		//se guarda el secret y los datos ingresados del usuario
		userStruct := usersStruct{&user, secret}

		//menu para decidir si se va a comparar los tokens
		for true{

			option := 0
			fmt.Println("1. Decodificar")
			fmt.Println("2. Salir:")

			fmt.Scanf("%d", &option)

			if option == 1 {

				confirmToken(jwt,&userStruct)
				fmt.Scanf("%d", &option)

			}
		}



	}


}