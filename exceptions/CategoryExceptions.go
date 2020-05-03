package exceptions

import "github.com/mohammadMghi/moduleForGo/internals/exception"

var (
	CategoryAlreadyExists = exception.Exception{
		Message: "category already exists",
		Status:  exception.AlreadyExists,
	}
	CategoryParentNotFound = exception.Exception{
		Message: "parent does not exist",
		Status:  exception.NotFound,
	}
)
