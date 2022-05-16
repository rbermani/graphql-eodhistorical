package main

import (
	"log"
	"net/http"
	"os"

	"gqlgen-cape/graph/generated"
	"gqlgen-cape/graph/resolver"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	tmpResolver, err := resolver.NewResolver()
	if err != nil {
		log.Printf("Unable to Create New Resolver and Initialize Db connection.")
		return
	}
	defer tmpResolver.Close()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: tmpResolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
