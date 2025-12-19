FROM golang:1.25

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

# Build with same flags as ORY Hydra
RUN go build -ldflags="-extldflags=-static" -o /usr/bin/test-crypto main.go

ENTRYPOINT ["/usr/bin/test-crypto"]
