package service

import (
	"encoding/json"
	"parameter-testing/domain"
	"parameter-testing/domain/constant"
	"parameter-testing/domain/entity"
)

func convertAPItoAPIResponse(api *entity.API) (*domain.APIResponse, error) {
	var extraData map[string]interface{}
	var fieldListRaw []domain.FieldList

	fieldList := make(map[string][]interface{})

	err := json.Unmarshal([]byte(api.ExtraData), &extraData)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(api.Field), &fieldListRaw)
	if err != nil {
		return nil, err
	}

	for _, field := range fieldListRaw {
		fieldList[field.In] = append(fieldList[field.In], field)
	}

	data := domain.APIResponse{
		ID:        api.ID,
		Path:      api.Path,
		Method:    constant.HTTP_METHOD_MAP[api.Method],
		Host:      api.Host,
		Scheme:    api.Scheme,
		IsActive:  api.IsActive,
		ExtraData: extraData,
		FieldList: fieldList,
	}

	return &data, nil
}
