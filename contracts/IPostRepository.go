package contracts

import "github.com/mohammadMghi/moduleForGo/entities"

type IPostRepository interface {
	List() ([]entities.Post, error)
	//Create(payload payloads.CategoryCreatePayload) (entities.Post, error)

}
