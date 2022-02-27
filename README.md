# ![Simple blog system](logo.png)
# blog-api
Simple Blog System.

## Quick start

Quick start you must install [docker](https://www.docker.com), [docker-compose](https://docs.docker.com/compose/), [Git](https://git-scm.com/) and [Fiber-Swagger](https://github.com/arsmn/fiber-swagger).

**Starts ready docker container**
```bash
mkdir database && chmod o+w ./database && chmod o+w ./tests/postman 
```

**Builds and tests**

```bash
git clone https://github.com/edlorenzo/blog-api.git
cd blog-api 
chmod a+x start.sh
./start.sh
```
Press <code>Ctrl + c</code> to stop application.

## Getting started

### Install Golang (go1.11+)

Please check the official golang installation guide before you start. [Official Documentation](https://golang.org/doc/install)
Also make sure you have installed go1.17+ version.

For more info and detailed instructions please check this guide: [Setting GOPATH](https://github.com/golang/go/wiki/SettingGOPATH)

### Initialize Swagger Docs
```bash
go install github.com/swaggo/swag/cmd/swag@latest
make swagger-init
```

### Working with makefile

```bash
make help
```

### Install dependencies

```bash
➜ go mod download
➜ make tidy
```

### Run

```bash
➜ go run main.go
```

### Build

```bash
➜ go build
```

### Tests

```bash
➜ make test
```
### Swagger UI
Access the application using the port `5007`.
Open url http://localhost:5007/swagger/index.html in browser.

### Authentication
Some API required to authenticate.
Once a user successfully logged in, simply copy the token id from the response to authorize, following this format.

```bash
Token eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDYwMzg1OTEsImlkIjoxfQ.P3XVfmA25WC9V6yvXpk5qoC-accIuHnIjKL_0rswVmg
```
# ![Simple blog system](token.png)