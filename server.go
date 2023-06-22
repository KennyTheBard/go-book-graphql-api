package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/KennyTheBard/go-books-graphql-api/db"
	"github.com/KennyTheBard/go-books-graphql-api/graph"
	"github.com/KennyTheBard/go-books-graphql-api/graph/dataloaders"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const defaultPort = "8080"

func main() {
	// setup logger
	config := zap.Config{
		Encoding:         "console",
		Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			TimeKey:     "time",
			EncodeLevel: zapcore.CapitalColorLevelEncoder,
			EncodeTime:  zapcore.ISO8601TimeEncoder,
		},
	}
	logger, err := config.Build()
	if err != nil {
		panic("Failed to initialize the logger: " + err.Error())
	}
	defer logger.Sync()

	// database connection
	db, err := db.InitDatabaseConnection()
	if err != nil {
		panic(err)
	}

	// configure graphql
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		DB:     db,
		Logger: logger,
	}}))
	loaders := dataloaders.NewLoaders(db)
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", dataloaders.Middleware(loaders, srv))

	// start server
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	logger.Info(fmt.Sprintf("Connect to http://localhost:%s/ for GraphQL playground", port))
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logger.Error(err.Error())
	}
}
