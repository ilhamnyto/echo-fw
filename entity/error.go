package entity

import "net/http"

type CustomError struct {
	Code           string      `json:"Code"`
	StatusCode     int         `json:"status_code"`
	Message        string      `json:"message"`
	AdditionalInfo interface{} `json:"additional_info,omitempty"`
}

var (
	generalError = CustomError{
		Code:       "GENERAL_ERROR",
		StatusCode: http.StatusInternalServerError,
		Message: "INTERNAL SERVER ERROR",
	}

	repositoryError = CustomError{
		Code: "REPOSITORY_ERROR",
		StatusCode: http.StatusInternalServerError,
		Message: "REPOSITORY ERROR",
	}

	notFoundError = CustomError{
		Code: "NOT_FOUND_ERROR",
		StatusCode: http.StatusNotFound,
		Message: "NOT FOUND",
	}

	unauthorizedError = CustomError{
		Code: "NOT_AUTHORIZED",
		StatusCode: http.StatusUnauthorized,
	}
)

func GeneralError() *CustomError {
	err := generalError
	return &err
}

func GeneralErrorWithAdditionalInfo(info interface{}) *CustomError {
	err := generalError
	err.AdditionalInfo = info
	return &err
}

func RepositoryError() *CustomError {
	err := repositoryError
	return &err
}

func RepositoryErrorWithAdditionalInfo(info interface{}) *CustomError {
	err := repositoryError
	err.AdditionalInfo = info
	return &err
}

func NotFoundError() *CustomError {
	err := notFoundError
	return &err
}

func NotFoundErrorWithAdditionalInfo(info interface{}) *CustomError {
	err := notFoundError
	err.AdditionalInfo = info
	return &err
}

func Unauthorized() *CustomError {
	err := unauthorizedError
	return &err
}

func UnauthorizedWithAdditionalInfo(info interface{}) *CustomError {
	err := unauthorizedError
	err.AdditionalInfo = info
	return &err
}
