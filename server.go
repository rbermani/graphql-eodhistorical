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

//func reqFundamentals(target interface{}) error {
//	var http_client = &http.Client{Timeout: 10 * time.Second}
//	r, err := http_client.Get(eod_fundamental_api + eod_api_key)
//	if err != nil {
//		return err
//	}
//	defer r.Body.Close()
//
//	return json.NewDecoder(r.Body).Decode(target)
//}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
