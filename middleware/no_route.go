package middleware

import (
	"fmt"
	"parameter-testing/handler"

	"github.com/gin-gonic/gin"
)

func NoRouteMiddleware(ctx *gin.Context) {
	handler.ResponseNotFound(ctx, fmt.Errorf("not found"))
}
