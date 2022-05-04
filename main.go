package main

import (
	"pchatserver/graph"
	"pchatserver/graph/generated"
	"pchatserver/router"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/labstack/echo/v4"
)

func main() {

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	e := router.NewRouter(echo.New(), srv)

	e.Logger.Fatal(e.Start(":8080"))
}
