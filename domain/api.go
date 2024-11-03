package domain

import (
	"github.com/gin-gonic/gin"
)

type APIRequest struct {
	Url              URL                    `json:"url" binding:"required"`
	MandatoryRequest map[string]interface{} `json:"mandatory_request,omitempty"`
	FieldList        []FieldList            `json:"field_list" binding:"required,dive"`
	Response         []Response             `json:"response" binding:"required,dive"`
}

type APIResponse struct {
	ID        int64                    `json:"id"`
	Path      string                   `json:"path"`
	Method    string                   `json:"method"`
	Host      string                   `json:"host"`
	Scheme    string                   `json:"scheme"`
	IsActive  int8                     `json:"is_active"`
	ExtraData map[string]interface{}   `json:"extra_data"`
	FieldList map[string][]interface{} `json:"field_list"`
}

type URL struct {
	Scheme string `json:"scheme" binding:"required"`
	Path   string `json:"path" binding:"required"`
	Host   string `json:"host" binding:"required"`
	Method int8   `json:"method" binding:"required"`
}

type FieldList struct {
	Name       string                 `json:"name" binding:"required"`
	Type       string                 `json:"type" binding:"required"`
	In         string                 `json:"in" binding:"required"`
	Mandatory  bool                   `json:"mandatory"`
	Properties []FieldList            `json:"properties,omitempty" binding:"dive"`
	ExtraData  map[string]interface{} `json:"extra_data,omitempty"`
}

type Response struct {
	Code       int                  `json:"code" binding:"required"`
	Properties []ResponseProperties `json:"properties,omitempty" binding:"dive"`
}

type ResponseProperties struct {
	Name  string      `json:"name" binding:"required"`
	Value interface{} `json:"value" binding:"required"`
}

func (v *APIRequest) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(v); err != nil {
		return err
	}

	return nil
}
