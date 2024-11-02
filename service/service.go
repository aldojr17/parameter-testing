package service

import (
	"parameter-testing/domain"
	"parameter-testing/repository"
	"parameter-testing/repository/cache"
	"parameter-testing/utils/pagination"
)

type Service struct {
	cache      *cache.Cache
	repository *repository.Repository
}

func NewService(
	cache *cache.Cache,
	repository *repository.Repository,
) *Service {
	return &Service{
		cache:      cache,
		repository: repository,
	}
}

func (s *Service) GetWithPagination(pageable pagination.Pageable) (*pagination.Page, error) {
	resp, err := s.repository.GetWithPagination(pageable)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *Service) Get() (*domain.StructName, error) {
	data, _ := s.cache.Get()
	if data != nil {
		return data, nil
	}

	resp, err := s.repository.GetByUUID("uuid")
	if err != nil {
		return nil, err
	}

	if err = s.cache.Set(resp); err != nil {
		return nil, err
	}

	return resp, nil
}
