package http

import (
	"github.com/labstack/echo/v4"
	"github.com/mohammadMghi/moduleForGo/contracts"
	"github.com/mohammadMghi/moduleForGo/drivers"
	"github.com/mohammadMghi/moduleForGo/internals/dependencies"
	"github.com/mohammadMghi/moduleForGo/payloads"

)

type applicationEndpoint struct {
	deps dependencies.CommonDependencies
	service contracts.IApplicationService
}

func (endpoint applicationEndpoint) Boot(transport interface{}) {
	t := transport.(*echo.Group)
	group := t.Group("/application")
	group.GET("/info", endpoint.info).Name = "Application.Info"
	group.GET("/ping", endpoint.ping).Name = "Application.Ping"
	group.GET("/echo", endpoint.echo).Name = "Application.Echo"
}

func NewApplicationEndpoint(deps dependencies.CommonDependencies, service contracts.IApplicationService) applicationEndpoint {
	return applicationEndpoint{deps,service }
}

func (endpoint applicationEndpoint) info(ctx echo.Context) error {
	payload, err := endpoint.service.Info()
	return drivers.PayloadToResponse(ctx, payload, err)
}

func (endpoint applicationEndpoint) ping(ctx echo.Context) error {
	payload, err := endpoint.service.Ping()
	return drivers.PayloadToResponse(ctx, payload, err)
}

func (endpoint applicationEndpoint) echo(ctx echo.Context) error {
	payload, err := drivers.RequestToPayload(ctx, new(payloads.MessagePayload))
	if err != nil {
		return err
	}
	validationErr := endpoint.deps.Validator.Validate(payload)
	return drivers.PayloadToResponse(ctx, payload, validationErr)
}