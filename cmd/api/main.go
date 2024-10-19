package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hhertout/graphql_api_boilerplate/graph"
	"github.com/hhertout/graphql_api_boilerplate/internal/resolvers"
	"github.com/hhertout/graphql_api_boilerplate/pkg/middleware"
	"go.uber.org/zap"
)

const defaultPort = "8080"
const BASE_URL = "/api"

func main() {
	logger, _ := zap.NewProduction()
	if os.Getenv("GO_ENV") == "development" {
		logger, _ = zap.NewDevelopment()
	}
	defer logger.Sync()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/api"))
	http.Handle(BASE_URL, middleware.Logger(srv, logger))

	if os.Getenv("GO_ENV") == "development" {
		logger.Sugar().Infof("🐹 Connect to http://localhost:%v/ for GraphQL playground", port)
		logger.Sugar().Infof("🐹 API available on http://localhost:%v%s", port, BASE_URL)
		logger.Sugar().Info("⚠️ Caution : The server will be running under development mode 🔨🔨")
	}

	logger.Sugar().Infof("🚀 Server lauch on port %v ✨", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
