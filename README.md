# go-money-api

Este projeto tem duas funções básicas, me ajudar nos estudos de golang e arquitetura de projetos e melhorar a qualidade do gerenciamento dos meus investimentos.

## startup

### run docker-compose
    docker-compose -f deployments/docker-compose.yml up
    
### run migrations
  
### run application
    go run cmd/go-money-api/main.go

##  tests
    go test ./...
with coverage

    go test -cover ./...

### coverage view

    go test -coverprofile=coverage.out ./...
    go tool cover -html=coverage.out

## Links uteis
- [project layout](https://github.com/golang-standards/project-layout)
- [tips for testing](https://medium.com/@matryer/5-simple-tips-and-tricks-for-writing-unit-tests-in-golang-619653f90742)
- [go-database-sql](http://go-database-sql.org/)