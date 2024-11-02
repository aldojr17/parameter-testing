package router

import (
	"parameter-testing/handler"
	"parameter-testing/initialize"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, app *initialize.Application) {
	handler := handler.NewHandler(app)
	router.GET("/models", handler.Get)
	router.POST("/generate", handler.Post)
	router.GET("/generation/:id", handler.Get)
}
