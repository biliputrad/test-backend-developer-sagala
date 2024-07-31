package main

import (
	"fmt"
	"log"
	routeRegister "test-backend-developer-sagala/api/route-register"
	"test-backend-developer-sagala/common/constants"
	"test-backend-developer-sagala/config/database/postgres"
	"test-backend-developer-sagala/config/env"
	"test-backend-developer-sagala/config/route"
)

func main() {
	//Init config app.env
	config, err := env.LoadConfig(".")
	if err != nil {
		message := fmt.Sprintf("%s can't load configuration file", constants.Configuration)
		log.Fatal(message)
	}

	//Init database
	db := postgres.ConfigurationPostgres(config)

	//Init router
	router := route.InitRouter(config)

	//Register routes
	apiV1 := router.Group("api/v1")
	routeRegister.RouteRegister(db, apiV1)

	//Run route
	route.RunRoute(config, router)
}
