# golang-ca-example

Implement golang projects with clean architecture
1. [How to implement clean architecture in golang](https://medium.com/@rayato159/how-to-implement-clean-architecture-in-golang-en-f50d66378ebf)

- Database setup
```bash
docker pull postgres:alpine

docker run --name cockroachdb -p 5432:5432 -e POSTGRES_PASSWORD=123456 -d postgres:alpine

docker exec -it cockroachdb bash

psql -U postgres

CREATE DATABASE cockroachdb;
```
- Server run
```bash
go run main.go
```

- Test endpoint
```bash
curl --location 'http://localhost:8080/v1/cockroach' \
--header 'Content-Type: application/json' \
--data '{
    "amount": 3
}'
```

---

2. [Building a CRUD app with Clean Architecture in Golang](https://dev.to/michinoins/building-a-crud-app-with-mysql-gorm-echo-and-clean-architecture-in-go-h6d)
