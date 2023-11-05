# Drivvn Backend Task
A simple API back-end task for prospective drivvn software developers.

### [Postman API Endpoints](https://www.postman.com/avionics-engineer-99532960/workspace/public-apis/collection/27922445-9a793100-b15a-45b0-9b87-ff99260f4ec2?action=share&creator=27922445)

## Prerequisites
- Docker

#### Without Docker
- Go 
- Node 
- npm
- Postgres

## Starting Up With Docker

**Create docker containers**
1. Go to root directory of project
2. Run `docker compose -p drivvn-josh up --build -d`

If backend service does not start for some reason
1. Run `docker compose -p drivvn-josh start` in the root directory

**Initialise database**
1. Run `docker exec -d drivvn-josh-backend-1 go run ./db/migrate/migrate.go`

You're now all set up! Visit http://localhost:8080/

## Starting Up Without Docker
Configure .env to point at your local db and make sure that the DB_URL are correct, e.g. host=localhost  dbname=drivvn etc...

1. Run `go mod tidy`
2. Run `go install`
3. Go to app directory, then
4. Run `npm install; npm run build`
5. Go to root of the project, then
6. Run: `go run db/migrate/migrate.go`
7. Run `go run main.go`
8. If the port hasn't been changed, visit http://localhost:8080/

## Running Test With Docker
1. Run `docker exec -it drivvn-josh-backend-1 bash`
1. Go to services `cd services`
2. Run `go test`

## Running Test Without Docker
1. Go to services `cd services`
2. Run `go test`

## Known bugs and limitations
- You cannot edit without providing BuildDate
- Frontend does not support error handling
- api test does not quite work yet but services do
- .env shouldn't be existing in multiple places, e.g. inside services and root directory