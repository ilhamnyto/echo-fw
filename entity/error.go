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
		Message: "INTERNAL SERVER ERROR",
		StatusCode: http.StatusInternalServerError,
	}

	repositoryError = CustomError{
		Code: "REPOSITORY_ERROR",
		Message: "REPOSITORY ERROR",
		StatusCode: http.StatusInternalServerError,
	}

	notFoundError = CustomError{
		Code: "NOT_FOUND_ERROR",
		Message: "NOT FOUND",
		StatusCode: http.StatusNotFound,
	}

	unauthorizedError = CustomError{
		Code: "NOT_AUTHORIZED",
		Message: "NOT AUTHORIZED REQUEST",
		StatusCode: http.StatusUnauthorized,
	}

	alreadyExistError = CustomError{
		Code: "ALREADY_EXIST",
		Message: "ALREADY EXIST DATA",
		StatusCode: http.StatusConflict,
	}

	badRequestError = CustomError{
		Code: "BAD_REQUEST",
		Message: "BAD REQUEST",
		StatusCode: http.StatusBadRequest,
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

func UnauthorizedError() *CustomError {
	err := unauthorizedError
	return &err
}

func UnauthorizedErrorWithAdditionalInfo(info interface{}) *CustomError {
	err := unauthorizedError
	err.AdditionalInfo = info
	return &err
}

func AlreadyExistError() *CustomError {
	err := alreadyExistError
	return &err
}

func AlreadyExistErrorWithAdditionalInfo(info interface{}) *CustomError {
	err := alreadyExistError
	err.AdditionalInfo = info
	return &err
}

func BadRequestError() *CustomError {
	err := badRequestError
	return &err
}

func BadRequestErrorWithAdditionalInfo(info interface{}) *CustomError {
	err := badRequestError
	err.AdditionalInfo = info
	return &err
}
