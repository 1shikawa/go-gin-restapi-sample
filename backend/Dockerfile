FROM golang:1.18-buster

COPY go.mod /go/src/
COPY go.sum /go/src/

WORKDIR /go/src

RUN go mod download

COPY . /go/src/

EXPOSE 3000

CMD [ "go", "run", "main.go" ]


