package handler

import (
	"parameter-testing/domain"
	"parameter-testing/initialize"
	"parameter-testing/repository"
	"parameter-testing/service"

	"github.com/gin-gonic/gin"
)

type APIHandler struct {
	apiService *service.APIService
}

func NewAPIHandler(app *initialize.Application) *APIHandler {
	return &APIHandler{
		apiService: service.NewAPIService(
			repository.NewAPIRepository(app.Database),
		),
	}
}

func (h *APIHandler) CreateAPI(c *gin.Context) {
	var payload domain.APIRequest

	if err := payload.Validate(c); err != nil {
		ResponseBadRequest(c, err)
		return
	}

	data, err := h.apiService.CreateAPI(payload)
	if err != nil {
		ResponseInternalServerError(c, err)
		return
	}

	ResponseCreated(c, data, "Successfully add API")
}
