package main

import (
	"fmt"
	"project1/app/middlewares"
	"project1/config"
	"project1/routers"
)

func main() {
	fmt.Println("Hello World")
	// Configuration to Database
	config.InitDB()
	// Call the router
	e := routers.New()
	middlewares.LogMiddlewares(e)
	// Logger to run server with port 8000
	e.Logger.Fatal(e.Start(":8000"))
}
