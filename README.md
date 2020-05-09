# num-converter

## Dev Environment Setup

### Local

Install Gin Web Framework

```sh
go get -u github.com/gin-gonic/gin
go get -u github.com/stretchr/testify/assert
go get -u github.com/sirupsen/logrus
```

## Run Docker

```sh
docker build -t num-converter .
docker run -p 8080:8080 num-converter
```
