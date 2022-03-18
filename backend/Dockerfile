FROM golang:1.18-buster as builder

COPY go.mod /go/src/
COPY go.sum /go/src/

WORKDIR /go/src

RUN go mod download

COPY . /go/src/

RUN go build -o app ./main.go


FROM alpine:latest
COPY --from=builder /go/src/app .

EXPOSE 3000

CMD [ "./app" ]
