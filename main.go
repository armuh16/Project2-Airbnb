package main

import (
	"alta/airbnb/config"
	"alta/airbnb/middlewares"
	"alta/airbnb/routers"
)

func main() {
	// Configuration to Database
	config.InitDB()
	// Call the router
	e := routers.New()
	middlewares.LogMiddlewares(e)
	// Logger to run server with port 8000
	e.Logger.Fatal(e.Start(":8000"))
}
