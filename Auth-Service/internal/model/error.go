package model

import (
	"errors"
	"github.com/hoangphuc28/CoursesOnline/Auth-Service/pkg/common"
	"net/http"
)

// Định nghĩa các error cho riêng phần User-Service
var (
	ErrEmailOrPasswordInvalid = common.NewCustomError(
		errors.New("email or password invalid"),
		http.StatusBadRequest,
		"email or password invalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		409,
		"email has already existed",
	)

	ErrCannotCreateAccount = common.NewCustomError(
		errors.New("can not create your account"),
		http.StatusBadRequest,
		"can not create your account",
	)
	ErrCannotUpdateUser = common.NewCustomError(
		errors.New("can not update your account"),
		http.StatusBadRequest,
		"can not update your account",
	)
	ErrAccountNotVerified = common.NewCustomError(
		errors.New("This account has not been verified!"),
		http.StatusForbidden,
		"This account has not been verified!")
)
