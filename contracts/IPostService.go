package contracts

import "github.com/mohammadMghi/moduleForGo/entities"

type IPostService interface {
	List() ([]entities.Post, error)

}
