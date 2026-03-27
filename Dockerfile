FROM golang:1.25.5

COPY config ./config

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o server ./cmd/server

EXPOSE 50051

CMD ["./server"]