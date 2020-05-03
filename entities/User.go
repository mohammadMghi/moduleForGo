package entities

import "github.com/mohammadMghi/moduleForGo/internals/entity"

type UserStatus int

const (
	PENDING UserStatus = iota
	ACTIVE
	BANNED
)

func (d UserStatus) String() string {
	return [...]string{"PENDING", "ACTIVE", "BANNED"}[d]
}

type User struct {
	entity.BaseEntity
	FirstName string
	LastName string
	Password string `json:"-"`
	Email string
	Status UserStatus
	Role string
}
