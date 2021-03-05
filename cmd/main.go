package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jim-at-jibba/gopher-notes/graph"
	"github.com/jim-at-jibba/gopher-notes/graph/generated"
	"github.com/jim-at-jibba/gopher-notes/pkg/repository"
	"github.com/jim-at-jibba/gopher-notes/pkg/service"
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

	dbPath := goDotEnvVariable("DB")
	// connect to db
	db, err := setupDatabase(dbPath)

	if err != nil {
		return err
	}

	// create new storage
	storage := repository.NewStorage(db)

	// run migrations
	err = storage.RunMigrations(dbPath)

	if err != nil {
		return err
	}

	// create new NotesService and pass storage in
	noteRepository := repository.NewNoteRepository(db)
	noteService := service.NewNoteService(noteRepository)

	// create user repo and service
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	// switch to chi router

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{NoteService: noteService, UserService: userService}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	return nil
}

func setupDatabase(connectionString string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", connectionString)

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
