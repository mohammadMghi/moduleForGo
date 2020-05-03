package contracts

import (
	"github.com/mohammadMghi/moduleForGo/entities"
	"github.com/mohammadMghi/moduleForGo/payloads"
)

type ICategoryService interface {
	List() ([]entities.Category, error)
	Create(payload payloads.CategoryCreatePayload) (entities.Category, error)
}
