FROM golang:1.20

WORKDIR /users-service

COPY . /users-service

EXPOSE 8080

RUN go mod download

RUN go build -o users-service .

CMD ["./users-service"]