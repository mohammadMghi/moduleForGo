package entities

import "github.com/mohammadMghi/moduleForGo/internals/entity"

type Setting struct {
	entity.BaseEntity
	Name  string
	Value string
}
