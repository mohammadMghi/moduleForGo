package contracts

import (
	"github.com/mohammadMghi/moduleForGo/entities"
	"github.com/mohammadMghi/moduleForGo/payloads"
)

type IUserService interface {
	Register(payload payloads.UserRegisterPayload) (entities.User, error)
	CheckCredentials(payload payloads.UserCredentialsPayload) (entities.User, error)
	Get(payload payloads.UserIDOnlyPayload) (entities.User, error)
}
