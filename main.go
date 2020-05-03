package main

import (
	"github.com/mohammadMghi/moduleForGo/drivers"
	"github.com/mohammadMghi/moduleForGo/endpoints/http"
	"github.com/mohammadMghi/moduleForGo/entities"
	"github.com/mohammadMghi/moduleForGo/internals/dependencies"
	"github.com/mohammadMghi/moduleForGo/repositories/mysql"
	"github.com/mohammadMghi/moduleForGo/services"
)


func main() {

	gateway := drivers.NewEchoGateway()
	config := drivers.NewViperConfiguration()
	validator := drivers.NewOzzoValidator()
	pool := drivers.NewGormDriver().Boot(config,entities.User{}, entities.Category{}, entities.Post{}, entities.Setting{})

	deps := dependencies.CommonDependencies{config,validator,}

	applicationRepository := mysql.NewApplicationRepository(deps, pool)
	applicationService := services.NewApplicationService(deps, applicationRepository)
	applicationEndpoint := http.NewApplicationEndpoint(deps,applicationService)

	userRepository := mysql.NewUserRepository(deps, pool)
	userService := services.NewUserService(deps,userRepository)
	userEndpoint := http.NewUserEndpoint(deps,userService)

	categoryRepository := mysql.NewCategoryRepository(deps, pool)
	categoryService := services.NewCategoryService(deps,categoryRepository)
	categoryEndpoint := http.NewCategoryEndpoint(deps,categoryService)

	postRepository := mysql.NewPostRepository(deps, pool)
	postService := services.NewPostService(deps, postRepository)
	postEndpoint := http.NewPostEndpoint(deps,postService)

	gateway.Boot(config, applicationEndpoint, userEndpoint, categoryEndpoint, postEndpoint)

}
