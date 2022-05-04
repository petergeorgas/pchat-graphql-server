package main

import (
	"log"
	"pchatserver/datastore"
	"pchatserver/graph"
	"pchatserver/graphql"
	"pchatserver/router"

	"github.com/labstack/echo/v4"
)

func main() {
	redisClient, err := datastore.NewRedisClient("localhost:6379") // TODO: FIX HARD-CODED addy
	if err != nil {
		log.Fatal("Failed to connect to Redis!")
	}
	defer redisClient.Close()

	resv := graph.NewResolver(redisClient)
	resv.SubscribeRedis()

	srv := graphql.NewGraphQLServer(resv)
	e := router.NewRouter(echo.New(), srv)

	e.Logger.Fatal(e.Start(":8080"))
}
