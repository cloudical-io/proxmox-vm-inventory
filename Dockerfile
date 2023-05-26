FROM golang:1.20

WORKDIR /build

COPY . .

RUN CGO_ENABLED=0 go mod download

RUN CGO_ENABLED=0 go build -o pveinventory cmd/main.go

ENTRYPOINT ["./pveinventory"]