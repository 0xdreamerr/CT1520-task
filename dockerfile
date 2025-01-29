FROM golang:alpine AS build

WORKDIR /url-shortener

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o shortener cmd/shortener/main.go

FROM alpine:latest

WORKDIR /url-shortener

COPY --from=build /url-shortener/shortener .

EXPOSE 8080

CMD ["./shortener", "-a=0.0.0.0:8080", "-b=gelasimov.center-dom.info:8080"]