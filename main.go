package main

import (
	"parameter-testing/initialize"
	"parameter-testing/middleware"
	"parameter-testing/router"

	"github.com/gin-gonic/gin"
)

func main() {
	app := initialize.InitApp()

	r := gin.New()
	r.NoRoute(middleware.NoRouteMiddleware)

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(middleware.APIMiddleware)

	router.Routes(r, app)

	if err := r.Run("localhost:8080"); err != nil {
		panic(err)
	}
}
