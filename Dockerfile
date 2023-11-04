FROM golang:1.21

WORKDIR /opt

COPY go.mod go.sum ./

RUN go mod tidy

EXPOSE 8080

COPY . .

CMD [ "go run main.go" ]
