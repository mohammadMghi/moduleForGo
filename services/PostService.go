package services

import (
	"github.com/mohammadMghi/moduleForGo/contracts"
	"github.com/mohammadMghi/moduleForGo/entities"
	"github.com/mohammadMghi/moduleForGo/internals/dependencies"
)

type postService struct {
	repository contracts.IPostRepository
}

func NewPostService(deps dependencies.CommonDependencies,repository contracts.IPostRepository) postService {
	return postService{
		repository,
	}
}

func (p postService) List() ([]entities.Post, error) {
	return p.repository.List()
}
