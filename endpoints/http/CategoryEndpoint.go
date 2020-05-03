package http

import (
	"github.com/labstack/echo/v4"
	"github.com/mohammadMghi/moduleForGo/contracts"
	"github.com/mohammadMghi/moduleForGo/drivers"
	"github.com/mohammadMghi/moduleForGo/internals/dependencies"
	"github.com/mohammadMghi/moduleForGo/payloads"
)

type categoryEndpoint struct {
	deps dependencies.CommonDependencies
	service contracts.ICategoryService
}

func (endpoint categoryEndpoint) Boot(transport interface{}) {
	t := transport.(*echo.Group)
	group := t.Group("/category")
	group.POST("/create", endpoint.create).Name = "Category.create"
	group.GET("/list", endpoint.list).Name = "Category.list"
}

func NewCategoryEndpoint(deps dependencies.CommonDependencies, service contracts.ICategoryService) categoryEndpoint {
	return categoryEndpoint{deps,service }
}

func (endpoint categoryEndpoint) list(ctx echo.Context) error {
	payload, err := endpoint.service.List()
	return drivers.TypeToResponse(ctx, payload, err)
}

func (endpoint categoryEndpoint) create(ctx echo.Context) error {
	input, err := drivers.RequestToPayload(ctx, new(payloads.CategoryCreatePayload))
	if err != nil {
		return err
	}
	payload, err := endpoint.service.Create(*input.(*payloads.CategoryCreatePayload))
	if err != nil {
		return err
	}
	return drivers.TypeToResponse(ctx, payload, err)
}
