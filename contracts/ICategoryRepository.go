package contracts

import "github.com/mohammadMghi/moduleForGo/entities"

type ICategoryRepository interface {
	Get(ID uint) (entities.Category, error)
	List() ([]entities.Category, error)
	Create(category entities.Category) (entities.Category, error)
	SlugExists(Slug string) (bool, error)
	Exists(ID uint) (bool, error)
}
