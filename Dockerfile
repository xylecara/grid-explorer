FROM golang:1.26.5 AS builder

WORKDIR /grid-explorer

COPY . .
RUN go build -o main .

FROM alpine AS final

WORKDIR /root/

COPY --from=builder /grid-explorer .

CMD ["./main"]
