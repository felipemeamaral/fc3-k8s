FROM golang:1.19.5-alpine as builder

WORKDIR /app

COPY go.mod server.go /app/

RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' -o server .

FROM scratch

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]