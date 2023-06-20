package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/KennyTheBard/go-books-graphql-api/db"
	"github.com/KennyTheBard/go-books-graphql-api/graph"
	"github.com/KennyTheBard/go-books-graphql-api/graph/dataloaders"
)

const defaultPort = "8080"

func main() {
	// database connection
	db, err := db.InitDatabaseConnection()
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		DB: db,
	}}))
	loaders := dataloaders.NewLoaders(db)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", dataloaders.Middleware(loaders, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
