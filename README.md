
# graphql-eodhistorical

## Description
This is a real world working example of a GraphQL server, written in Golang, that consolidates multiple RESTful API endpoints into a unified GraphQL interface.
The example uses the EOD Historical Data API <http://https://eodhistoricaldata.com/>

Each time a request is made, the result is cached in a MongoDB collection. This prevents redundant server requests in the same day.

## Typical Usage

### Re-Generate the models:

```
$ go run github.com/99designs/gqlgen generate
```

### Initialize Docker MongoDB instance

```
$ sudo systemctl start docker
$ cd mongodb
$ sudo docker compose up -d
$ sudo docker exec -it mongodb bash
```

### Start Server
```
$ cd .. ; go build
$ go run server.go
```

Connect to the localhost GraphQL playground at port :8080


