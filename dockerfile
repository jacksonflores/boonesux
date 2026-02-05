FROM golang:latest-alpine as build

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY  . .

RUN CGO_ENABLED=0 GOOS=linux go build -o boonesux ./cmd

FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=build /app/boonesux /app/boonesux

COPY web /app/web

EXPOSE 8080

CMD ["/app/boonesux"]
