# Building Distributed Applications in Gin

Following the book for learn to create apps in GO.

### Note

If you're running Go 1.16 and above, you need to disable Go modules via the

```bash
GO111MODULE=off option
```

### Note

If you're working with a team of developers, you will need to issue the `go
mod download` command to install the required dependencies after cloning
the project from GitHub.

## Go Swagger

Install [go-swagger](https://github.com/go-swagger/go-swagger)

1. Generate swagger.json / OpenAPI V2

```bash
go-swagger generate spec -o ./swagger.json
```

2. Convert swagger.json to openapi.json / OpenAPI V2 to V3

### Manual

Go to [swagger-convert](https://converter.swagger.io/) and convert the swagger.json to openapi.json.

Download the ***response***.

### Auto

```bash
curl -X POST "https://converter.swagger.io/api/convert" -H "accept: application/json" -H "Content-Type: application/json" -d "@./swagger.json" > openapi.json
```

3. Docker openapi image

```bash
docker pull swaggerapi/swagger-ui
```

```bash
docker run --rm -p 80:8080 -e SWAGGER_JSON=/app/openapi.json -v @path/to/golang/app:/app swaggerapi/swagger-ui
```

### Test swagger-ui

```go
go run main.go
```

### Swagger UI

[URL Localhost](http:localhost:80)
