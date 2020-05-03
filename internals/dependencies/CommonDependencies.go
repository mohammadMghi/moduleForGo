package dependencies

import (
	contracts2 "github.com/mohammadMghi/moduleForGo/internals/contracts"
)

type CommonDependencies struct {
	Configuration contracts2.IConfiguration
	Validator     contracts2.IValidator
}