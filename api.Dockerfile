
FROM golang:1.19-alpine

WORKDIR /balance_avito

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o /build  cmd/service/main.go

EXPOSE 1323

CMD [ "/build" ]
