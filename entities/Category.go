package entities

import "github.com/mohammadMghi/moduleForGo/internals/entity"

type Category struct {
	entity.BaseEntity
	Title string
	Slug string
	ParentID uint
	Parent *Category
}

