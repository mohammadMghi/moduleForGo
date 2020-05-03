package mysql

import (
	"github.com/mohammadMghi/moduleForGo/drivers"
	"github.com/mohammadMghi/moduleForGo/entities"
	"github.com/mohammadMghi/moduleForGo/internals/dependencies"
	"testing"
)

func TestApplicationRepositoryInfoPayload(t *testing.T) {
	deps := dependencies.CommonDependencies{
		Configuration: drivers.NewViperConfiguration(),
		Validator:     drivers.NewOzzoValidator(),
	}
	repo := NewApplicationRepository(deps ,drivers.NewGormDriver().Boot(deps.Configuration,entities.User{}))
	payload, err := repo.Info()
	if err != nil {
		t.Fatalf("%v" ,err.Error())
	}
	err = deps.Validator.Validate(payload)

	if err != nil {
		t.Fatalf("%v" ,err.Error())
	}
}
