##### Vídeo:

https://www.youtube.com/watch?v=7UwdiSdRmQA



##### Docker file utilizado para dockerizar los servicios de go

```dockerfile
FROM golang:latest
RUN apt-get update
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64


WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o main .
WORKDIR /dist
RUN cp /build/main .
EXPOSE 8080
CMD ["/dist/main"]
```



##### Docker-compose con los servicios de la practica 7

```yaml
version: '3'

services:

  esb:
    build: esb/.
    ports:
        - 8085:8085
    networks:
      service_network:


  cliente:
    build: serivicio-cliente/.
    ports:
      - 8080:8080
    networks:
      service_network:
    depends_on:
      - esb

  repartidor:
    build: servicio-repartidor/.
    ports:
      - 8082:8082
    networks:
        service_network:
    depends_on:
      - esb


  restaurante:
    build: servicio-restaruante/.
    ports:
      - 8081:8081
    networks:
      service_network:
    depends_on:
      - esb


networks:
    service_network:
        driver: bridge

    default:
      external:
        name: service_network
```

