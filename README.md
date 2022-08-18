# Building Distributed Applications in Gin

Following the book for learn to create apps in GO.

## Note
If you're running Go 1.16 and above, you need to disable Go modules via the
```bash
GO111MODULE=off option
```

## Note
If you're working with a team of developers, you will need to issue the `go
mod download` command to install the required dependencies after cloning
the project from GitHub.

## Note Swagger GO

### Docker image

```bash
docker pull quay.io/goswagger/swagger
```

### For Mac and Linux users

```bash
alias swagger='docker run --rm -it  --user $(id -u):$(id -g) -e GOPATH=$(go env GOPATH):/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger'
swagger version
```
