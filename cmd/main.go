package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jim-at-jibba/gopher-notes/graph"
	"github.com/jim-at-jibba/gopher-notes/graph/generated"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	run()
}

// refactor to run functions
func run() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// connect to db

	// create new storage

	// run migrations

	// create new NotesService and pass storage in

	// switch to chi router

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
