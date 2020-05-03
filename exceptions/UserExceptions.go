package exceptions

import "github.com/mohammadMghi/moduleForGo/internals/exception"

var (
	UserNotFound = exception.Exception{
		Message: "user not found",
		Status:  exception.NotFound,
	}
	UserAlreadyExists = exception.Exception{
		Message: "user already exists",
		Status:  exception.AlreadyExists,
	}

	InvalidCredentials = exception.Exception{
		Message: "invalid credentials",
		Status:  exception.InvalidInput,
	}
)
