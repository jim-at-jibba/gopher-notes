package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jim-at-jibba/gopher-notes/graph"
	"github.com/jim-at-jibba/gopher-notes/graph/generated"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\n", err)
		os.Exit(1)
	}
}

// refactor to run functions
func run() error {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// connect to db
	_, err := setupDatabase()

	if err != nil {
		return err
	}

	// create new storage

	// run migrations

	// create new NotesService and pass storage in

	// switch to chi router

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	return nil
}

func setupDatabase() (*sqlx.DB, error) {
	dbPath := goDotEnvVariable("DB")
	db, err := sqlx.Open("postgres", dbPath)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}

func goDotEnvVariable(key string) string {

	err := godotenv.Load("../.env")

	if err != nil {

		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)

}
