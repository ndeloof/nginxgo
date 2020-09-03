FROM golang:1.14 AS build

RUN apt-get update && apt-get install -y curl
EXPOSE 80
HEALTHCHECK --interval=10s \
  CMD curl -f http://localhost:80
  
WORKDIR /compose/hello-docker
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o backend main.go
CMD ["/compose/hello-docker/backend"]

