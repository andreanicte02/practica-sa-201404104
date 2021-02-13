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

type user struct {

	Name string
	Carne int

}


type header struct {

	Alg string `json:"alg"`
	Typ string `json:"type"`
}

type usersStruct struct {

	user *user
	secret string

}

func base64Encode(src string) string {
	return strings.
		TrimRight(base64.URLEncoding.
			EncodeToString([]byte(src)), "=")
}


func base64Decode(src string) string {
	if l:= len(src) % 4; l > 0 {
		src += strings.Repeat("=", 4-l)
	}
	decoded, _ := base64.URLEncoding.DecodeString(src)

	return string(decoded)
}

func generateScret() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}




func main() {

	menu()


}

func generateSignature(secret string, strFinal string) string {

	key:= []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(strFinal))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))

	
}

func generateJWT(user * user, secret string) string  {

	headerData := header {"HS256", "JWT"}
	jsonHeader, _ := json.Marshal(headerData)

	encodedHeader:= base64Encode(string(jsonHeader))

	jsonPayload,_ := json.Marshal(user)
	encodedPyload := base64Encode(string(jsonPayload))

	strFinal := encodedHeader+"."+encodedPyload

	return strFinal+"."+generateSignature(secret,strFinal)



}

func confirmToken(jwt string, userStruct *usersStruct) {
	arrayJWT := strings.Split(jwt, ".")

	if len(arrayJWT) != 3{
		fmt.Println("el token no esta completo")
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









func menu(){


	for true {

		var user user
		secret := generateScret()
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println(">>>>>>>>>>>>>>>>>>>>menu<<<<<<<<<<<<<<<<<<<<<<")
		fmt.Println("Ingresar nombre:")

		scanner.Scan()
		user.Name = scanner.Text()

		fmt.Println("Carnet:")

		scanner.Scan()
		user.Carne, _ = strconv.Atoi(scanner.Text())

		jwt :=generateJWT(&user, secret)
		fmt.Println("Token generado: "+ jwt)
		userStruct := usersStruct{&user, secret}

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