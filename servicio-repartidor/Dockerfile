FROM golang:latest
RUN apt-get update
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64


WORKDIR /build
COPY practica-sa-201404104/serivicio-cliente/go.mod .
COPY practica-sa-201404104/serivicio-cliente/go.sum .
RUN go mod download

COPY practica-sa-201404104/serivicio-cliente/. .
RUN go build -o main .
WORKDIR /dist
RUN cp /build/main .
EXPOSE 8080
CMD ["/dist/main"]