package estructura


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



