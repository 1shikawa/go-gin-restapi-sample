FROM golang:1.18-buster as builder

WORKDIR /go/src

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /app main.go


FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /app /app

EXPOSE 3000

USER nonroot:nonroot

ENTRYPOINT [ "/app" ]
