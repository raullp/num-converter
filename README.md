# num-converter

## Dev Environment Setup

### Local

num-converter uses [GoLang modules](https://blog.golang.org/using-go-modules) which is on charge of dowloading the dependencies.

```sh
go build
go run main.go
```

## Run Docker

```sh
docker build -t num-converter .
docker run -p 8080:8080 num-converter
```
