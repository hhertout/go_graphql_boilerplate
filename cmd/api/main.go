package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hhertout/graphql_api_boilerplate/graph"
	"github.com/hhertout/graphql_api_boilerplate/internal/application/middleware"
	"github.com/hhertout/graphql_api_boilerplate/internal/application/resolvers"
	"go.uber.org/zap"
)

const defaultPort = "4000"
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

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{}}))
	srv.AddTransport(transport.POST{})
	srv.Use(extension.Introspection{})

	srv.AroundOperations(middleware.AddLoggerToContext(logger))
	srv.AroundOperations(middleware.Logger)

	http.Handle("/", playground.Handler("GraphQL playground", "/api"))
	http.Handle(BASE_URL, srv)

	if os.Getenv("GO_ENV") == "development" {
		logger.Sugar().Infof("🐹 Connect to http://localhost:%v/ for GraphQL playground", port)
		logger.Sugar().Infof("🐹 API available on http://localhost:%v%s", port, BASE_URL)
		logger.Sugar().Info("⚠️ Caution : The server will be running under development mode 🔨🔨")
	}

	logger.Sugar().Infof("🚀 Server lauch on port %v ✨", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
