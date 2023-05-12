package utils

import (
	"strings"

	"github.com/ilhamnyto/echo-fw/entity"
)

func ValidateRegisterRequest(req *entity.CreateUserRequest) *entity.CustomError {
	if !strings.Contains(req.Email, "@") {
		return entity.BadRequestErrorWithAdditionalInfo("Wrong email format")
	}
	if len(req.Username) < 6 {
		return entity.BadRequestErrorWithAdditionalInfo("Username should have at least 6 characters")
	}
	if len(req.Password) < 8 {
		return entity.BadRequestErrorWithAdditionalInfo("Password should have at least 8 characters")
	}
	return nil
}