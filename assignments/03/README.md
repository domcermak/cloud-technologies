
# Go REST Api

### Assignment

Product management API

1. Create golang server app that provides API for managing list of products
   1. Product(name, price, amount)
   2. Operations - get, list, update, delete
   4. Backend - SQL or NOSQL database (Redis, Mongo, Cassandra)
   3. Include simple testing client in your repository

## Run server
> docker-compose up

## Run client
> go run client.go

## Run tests
### First terminal window
> docker-compose -f tests/db/docker-compose.yml up

### Second terminal window
> go test -tags "integration" ./tests/client ./tests/server ./tests/integration