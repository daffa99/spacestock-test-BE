package main

import (
	m "BE-SpaceStock-Test/Assignment-7/middlewares"
	"BE-SpaceStock-Test/Assignment-7/routes"
)

func main() {
	e := routes.New()
	// To show logs
	m.LogMiddleware(e)
	address := ":3000"
	// Start server and logs
	e.Logger.Fatal(e.Start(address))
}
