package router

import (
	"parameter-testing/handler"
	"parameter-testing/initialize"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, app *initialize.Application) {
	apiHandler := handler.NewAPIHandler(app)

	parameterTestingGroup := router.Group("/v1/parameter_testing")
	{
		apiGroup := parameterTestingGroup.Group("/api")
		{
			apiGroup.POST("", apiHandler.CreateAPI)
		}
	}
}
