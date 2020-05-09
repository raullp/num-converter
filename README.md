# num-converter

## Tech Notes

### Architecture

NumConverter is a simple stateless Rest API that can easily scale horizontally scale. The API layer is built on top of [Gin](https://github.com/gin-gonic/gin) which has a good performance with a minimalistic design.

### Security

NumConverter microservice assumes that the Authorization filters are done in the top layer. By the way, Gin has the capability to add an interceptor to either implement the Security internally or reach out to an authorization microservice.

### Error Handling

A centralized place for error handling was implemented, with the ability to convert errors to http response errors (Gin interceptor), provide useful information about caller context, and a centralized place for logging SystemErrors.

```go
    error.E(errors.New("number: is not a valid number"), 'converter.numToWords'KindBadRequest, log.ErrorLevel)
```

```go
    SystemError(err)
```

### Logging

Logging is based on [logrus](https://github.com/sirupsen/logrus) which has the ability to export logs as JSON objects which can be hooked to ELK, Watchdog, Splunk, etc.

### Dockerized

NumConverter microservice has been dockerized, so it is easy to be used by developers or deploy to containerized platforms like Kubernetes.

### API Documentation

Developers can find documentation of its API on http://domain:8080/swagger/index.html . The documentation has been generated from code using [Swaggo](https://github.com/swaggo/swag) which generates a Swagger file.

## Dev Environment Setup

### Local

num-converter uses [GoLang modules](https://blog.golang.org/using-go-modules) which is on the charge of downloading the dependencies.

#### Install swag

Install swag to generate the docs, and run swag init every time you update your doc files

```sh
go get -u github.com/swaggo/swag/cmd/swag
~/go/bin/swag init
```

```sh
go build
go run main.go
```

### Run Docker

```sh
docker build -t num-converter .
docker run -p 8080:8080 num-converter
```
