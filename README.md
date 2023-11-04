# Drivvn Backend Task

## Prerequisites
- Go 
- Node 
- npm
- Postgres (Check Postgres Docker section)

## Starting Up
Configure .env to point at your local db

1. Run `go mod tidy`
2. Run `go install`
3. Go to app directory, then
4. Run `npm install; npm run build`
5. Go to root of the project, then
6. Run: `go run db/migrate/migrate.go`
7. Run `go run main.go`
8. If the port hasn't been changed, visit http://localhost:8080/

## Running Test
1. Go to services `cd services`
2. Run `go test`

## Postgres Docker
If you don't have Postgres installed but you have Docker then do the following:

#### Start Postgres Instance
Run `docker run --name postgres -e POSTGRES_PASSWORD=root -e POSTGRES_USER=root -p 5432:5432 -d postgres`

#### Create DB
Run `docker exec -it postgres createdb --username=root --owner=root drivvn`
