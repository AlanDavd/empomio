# Empomio

Retail and wholesale marketplace for the clothing industry in Mexico.

## Set Up

Install go
Setup private repositories for Golang

## Run

Docker for local development

```bash
docker-compose up
```

Run server

```bash
# Without flags
go run cmd/http/main.go

# With flags
go run cmd/http/main.go -address=:8081
```
