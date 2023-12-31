FROM golang:1.19

WORKDIR /app

COPY . .

RUN go mod download

CMD ["go", "test", "-bench=.", "-benchmem", "-benchtime=100x"]
