package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/mohammadMghi/moduleForGo/entities"
	"github.com/mohammadMghi/moduleForGo/internals/dependencies"
)

type postRepository struct {
	pool *gorm.DB
	deps dependencies.CommonDependencies
}

func (p postRepository) List() ([]entities.Post, error) {
	panic("implement me")
}

func NewPostRepository(deps dependencies.CommonDependencies, pool interface{}) postRepository {
	return postRepository{pool: pool.(*gorm.DB), deps: deps}
}
