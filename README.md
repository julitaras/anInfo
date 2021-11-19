# Go Starter App

![technology Go](https://img.shields.io/badge/technology-go-blue.svg)

This is a basic Go application

# Instalation

## Previous requirements:

1. Install Go: https://golang.og/doc/install

2. Create `.env` file in the same root where `main.go` file is, with:

```env
PORT = "8080"
DB_USERNAME = "user"
DB_PASSWORD = "pass"
DB_HOST = "host.com"
DB_NAME = "dbName"
GIN_MODE = "debug"
```

## Execution

```sh
cd src/api
go run .
```
