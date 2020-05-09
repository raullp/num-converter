FROM golang:1.13-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build

EXPOSE 8080
ENTRYPOINT [ "/app/num-converter" ]
