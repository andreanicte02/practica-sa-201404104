

Andrea Nicte Vicente Campos

201404104

# Practica 3

Link funcionalidad: https://youtu.be/92HZeT6-V3A

#### Structs que ayudaron a mantener la data en memoria:

```golang
//struct que nos va ayudar a simular los menus en memoria
type Menu struct {
	Id int
	Descripcion string
}

//struct que nos va ayudar a simular los clientes en memoria
type Cliente struct {
	Id int
	Nombre string
}

//struct que va ayudar a almacenar la informacion de pedidos en memoria del restaurante
type Pedido struct {

	IdMenu int `json:"idMenu"`
	IdCliente int `json:"idCliente"`
	IdEstado int `json:"IdEstado"` //0 pendiente 1 completado


}


//struct que va ayudar a almacenar la informacion de pedidos en memoria del repartidor
type PedidoRepartidor struct {

	IdMenu int `json:"idMenu"`
	IdCliente int `json:"idCliente"`
	IdEstado int `json:"IdEstado"` //0 pendiente 1 completado
	DescripcionMenu string `json:"DescripcionMenu"`
	IdPedido int  `json:"idPedido"`
	EstadoRepartidor int `json:"estadoRepartidor ya tomo el pedido o no"`


}
```

### Structs que ayudaron a enviar información en formato JSON de un servicio a otro:

```go
//struct de un json generico
type JSONGenerico struct {

	Id int `json:"id"`

}
//struct de un mensaje generico
type JSONMessageGeneric struct {

	Message string
	Id      int

}
```



### Función que permite manejar los endpoints de los servicios:

Restaurante

```go
func handle()  {

	router := mux.NewRouter()
	router.HandleFunc("/recibir_pedido",recibir_pedidio).Methods("POST")
	router.HandleFunc("/estado_pedido",etado_pedido).Methods("GET")
	router.HandleFunc("/avisar_pedido_listo",avisar_pedido_listo).Methods("POST")
	http.ListenAndServe(":8081", router)

}
```

Reapartidor

```
func handle()  {

	router := mux.NewRouter()
	router.HandleFunc("/recibir_pedidio",recibir_pedido).Methods("POST")
	router.HandleFunc("/informar_estado_cliente",informar_estado_cliente).Methods("GET")
	router.HandleFunc("/marcar_pedido",marcar_pedido).Methods("POST")
	http.ListenAndServe(":8082", router)

}
```



## Restaurante:

Restaurante cuenta con tres servicios:

- recibir_pedido
- avisar_pedido_listo
- estado_pedido



Ejemplo:

Esta función espera una solicitud por medio de 'POST', se utiliza para recibir los pedidos y guardar en memoria

```GO
func recibir_pedidio(w http.ResponseWriter, r *http.Request)  {

	data:= estructura.Pedido{}
	m:= estructura.JSONMessageGeneric{}
	w.Header().Set("Content-Type","application/json")

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_,existeCliente := hashCliente[data.IdCliente]
	_,existeMenu := hashMenu[data.IdMenu]

	if !existeCliente || !existeMenu{

		mensaje_error,_ := json.Marshal(estructura.JSONMessageGeneric{"El menu o el cliente no existen",-1})
		http.Error(w, string(mensaje_error), http.StatusBadRequest)
		return
	}


	//se retorna un json con el id del pedido
	hashPedido[idPedido]=&data
	m.Message =  "pedidio realizado"
	m.Id = idPedido
    //se escribe el json que se va enviar
	json.NewEncoder(w).Encode(m)
	fmt.Print("data recibida: ")
	fmt.Println(data)
	idPedido ++


}
```



## Repartidor:



Cuenta con 3 servicios:

- recibir_pedido
- informar_estado_cliente
- marcar_pedido

Ejemplo:

```
/**
	w -> indica el contenido de la respuesta
	r -> indica el contenido de la solicitud
*/

```



```go
//servicio que recibe el pedido, indica al reapartidor que tiene pedidos que entregar
//recibe una json con la estrucutra de PedidoRepartidor y actualiza el estado
func recibir_pedido(w http.ResponseWriter, r *http.Request)  {

	data:= estructura.PedidoRepartidor{}

	m:= estructura.JSONMessageGeneric{}
	w.Header().Set("Content-Type","application/json")

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if estadoRepartidor == 0 {
		data.EstadoRepartidor = 1
	}
	hashPedido[data.IdPedido]=&data

	m.Message =  "pedidio realizado"
	m.Id = data.IdPedido
	json.NewEncoder(w).Encode(m)
	fmt.Print("pedido recibida: ")
	fmt.Println(data)

	if estadoRepartidor == 1{
		return
	}
	estadoRepartidor = 1
}
```







### Cliente

Funciona como un cliente, envía solicitudes a las dos servidores , solicitudes:

- solicitar_estado_restaurante
- estado_repartidor
- solicitar_pedido

Ejemplo:

```
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
```



```go
func solicitar_pedido(pedido *estructura.Pedido){
	
	dataRequest,_:= json.Marshal(pedido)
	req,err := http.NewRequest("POST", "http://localhost:8081/recibir_pedido", bytes.NewBuffer(dataRequest))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	defer resp.Body.Close()

	var data estructura.JSONMessageGeneric
	decoder:= json.NewDecoder(resp.Body)
	decoder.Decode(&data)

	hashPedido[data.Id]=data.Id
	idPedido = data.Id
	fmt.Println("...respuesta")
	fmt.Println(data)
	fmt.Println("id del pedido:")
	fmt.Println(data.Id)


}
```

