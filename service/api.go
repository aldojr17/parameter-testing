package service

import (
	"encoding/json"
	"parameter-testing/domain"
	"parameter-testing/domain/entity"
	"parameter-testing/repository"
	"parameter-testing/util"
)

type APIService struct {
	apiRepo *repository.APIRepository
}

func NewAPIService(
	apiRepo *repository.APIRepository,
) *APIService {
	return &APIService{
		apiRepo: apiRepo,
	}
}

func (s *APIService) CreateAPI(payload domain.APIRequest) (*domain.APIResponse, error) {
	field, err := json.Marshal(payload.FieldList)
	if err != nil {
		return nil, err
	}

	extraDataRaw := map[string]interface{}{
		"response": payload.Response,
	}

	if len(payload.MandatoryRequest) != 0 {
		extraDataRaw["mandatory_request"] = payload.MandatoryRequest
	}

	extraData, err := json.Marshal(extraDataRaw)
	if err != nil {
		return nil, err
	}

	request := entity.API{
		Path:       payload.Url.Path,
		Method:     payload.Url.Method,
		Host:       payload.Url.Host,
		Scheme:     payload.Url.Scheme,
		Field:      string(field),
		ExtraData:  string(extraData),
		CreateTime: util.GenerateCurrentTimestamp(),
		UpdateTime: util.GenerateCurrentTimestamp(),
	}

	err = s.apiRepo.Create(&request)
	if err != nil {
		return nil, err
	}

	data, err := convertAPItoAPIResponse(&request)
	if err != nil {
		return nil, err
	}

	return data, nil
}
