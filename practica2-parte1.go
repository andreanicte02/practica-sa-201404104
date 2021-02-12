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

type secret struct {

	userData* user
	secret string
}

type header struct {

	Alg string `json:"alg"`
	Typ string `json:"type"`
}

func Base64Encode(src string) string {
	return strings.
		TrimRight(base64.URLEncoding.
			EncodeToString([]byte(src)), "=")
}


func Base64Decode(src string) string {
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

func generateJWS(userData * user) string  {

	headerData := header {"HS256", "JWT"}
	jsonHeader, _ := json.Marshal(headerData)

	encodedHeader:= Base64Encode(string(jsonHeader))

	jsonPayload,_ := json.Marshal(userData)
	encodedPyload := Base64Encode(string(jsonPayload))

	strFinal := encodedHeader+"."+encodedPyload

	return strFinal+"."+generateSignature(generateScret(),strFinal)



}




func menu(){

	var userData user
	var sercte := generateScret()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>menu<<<<<<<<<<<<<<<<<<<<<<")
	fmt.Println("Ingresar nombre:")

	scanner.Scan()
	userData.Name = scanner.Text()

	fmt.Println("Carnet:")

	scanner.Scan()
	userData.Carne, _ = strconv.Atoi(scanner.Text())

	fmt.Println(generateJWS(&userData))


}