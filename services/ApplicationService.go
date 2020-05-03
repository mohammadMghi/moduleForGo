package services

import (
	"github.com/mohammadMghi/moduleForGo/contracts"
	"github.com/mohammadMghi/moduleForGo/internals/dependencies"
	"github.com/mohammadMghi/moduleForGo/payloads"
)

type applicationService struct {
	repository contracts.IApplicationRepository
}

func NewApplicationService(deps dependencies.CommonDependencies,repository contracts.IApplicationRepository) applicationService {
	return applicationService{
		repository,
	}
}

func (service applicationService) Info() (payloads.ApplicationInfoPayload, error) {
	return service.repository.Info()
}

func (service applicationService) Ping() (payloads.ApplicationPingPayload, error) {
	return service.repository.Ping()
}
