FROM golang:1.23.1-alpine as build

WORKDIR /app

COPY ../ ./

COPY ./server/src/template ../template

RUN  go mod tidy

WORKDIR /app/server/src

RUN go build -o /backend

RUN rm -rf /app

WORKDIR /

EXPOSE 8080

CMD ["/backend"]
