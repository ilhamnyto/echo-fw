package entity

import "net/http"

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Payload    interface{} `json:"payload,omitempty"`
}

var (
	generalSuccess = Response{
		StatusCode: http.StatusOK,
		Message: "SUCCESS",
	}

	createSuccess = Response{
		StatusCode: http.StatusCreated,
		Message: "CREATED_SUCCESS",
	}
)

func GeneralSuccess() *Response {
	succ := generalSuccess
	return &succ
}

func GeneralSuccessWithCustomMessageAndPayload(message string, payload interface{}) *Response {
	succ := generalSuccess
	succ.Message = message
	succ.Payload = payload
	return &succ
}

func CreatedSuccess() *Response {
	succ := createSuccess
	return &succ
}

func CreatedSuccessWithPayload(paylaod interface{}) *Response {
	succ := createSuccess
	createSuccess.Payload = paylaod
	return &succ
}