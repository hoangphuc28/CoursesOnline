package common

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Causes     error  `json:"-"`
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
	// Dùng làm đa ngôn ngữ
	// Key    string `json:"key"`
}

func NewErrorResponse(causes error, status int, msg string) *AppError {
	return &AppError{causes, status, msg}
}
func (e *AppError) RootCauses() error {
	if err, ok := e.Causes.(*AppError); ok {
		return err.RootCauses()
	}
	return e.Causes
}

func (e *AppError) Error() string {
	return e.RootCauses().Error()
}

func ErrUnauthorized(causes error) *AppError {
	return NewErrorResponse(causes, http.StatusUnauthorized, causes.Error())
}

func ErrForbidden(causes error) *AppError {
	return NewErrorResponse(causes, http.StatusForbidden, "you have no permission")
}

func ErrBadRequest(causes error) *AppError {
	return NewErrorResponse(causes, http.StatusBadRequest, "invalid request")
}

func ErrNotFound(causes error) *AppError {
	return &AppError{causes, http.StatusNotFound, "not found"}
}

func ErrDB(causes error) *AppError {
	return &AppError{causes, http.StatusInternalServerError, "something went wrong with Database"}
}

func ErrInternal(causes error) *AppError {
	return &AppError{causes, http.StatusInternalServerError, "something went wrong in the server"}
}

func NewCustomError(causes error, statusCode int, msg string) *AppError {
	return &AppError{causes, statusCode, msg}
}

func ErrCannotListEntity(entity string, causes error) *AppError {
	return NewCustomError(causes, http.StatusBadRequest, fmt.Sprintf("Cannot list %s", entity))
}

func ErrEntityNotFound(entity string, causes error) *AppError {
	return NewCustomError(causes, http.StatusBadRequest, fmt.Sprintf("%s not found", entity))
}

func ErrCannotCreateEntity(entity string, causes error) *AppError {
	return NewCustomError(causes, http.StatusBadRequest, fmt.Sprintf("cannot create %s", entity))
}

func ErrCannotUpdateEntity(entity string, causes error) *AppError {
	return NewCustomError(causes, http.StatusBadRequest, fmt.Sprintf("cannot update %s", entity))
}

func ErrCannotDeleteEntity(entity string, causes error) *AppError {
	return NewCustomError(causes, http.StatusBadRequest, fmt.Sprintf("cannot delete %s", entity))
}

func ErrEntityExisted(entity string, causes error) *AppError {
	return NewCustomError(causes, http.StatusBadRequest, fmt.Sprintf("%s already exists", entity))
}
