FROM golang:1.25-alpine AS build

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY  . .

RUN CGO_ENABLED=0 GOOS=linux go build -o boonesux ./cmd

FROM alpine:3.20

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=build /app/boonesux /app/boonesux

COPY web /app/web

EXPOSE 8080

CMD ["/app/boonesux"]
