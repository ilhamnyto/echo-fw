package utils

import (
	"errors"
	"regexp"
	"strings"

	"github.com/ilhamnyto/echo-fw/entity"
)

func ValidateRegisterRequest(req *entity.CreateUserRequest) error {
	if !strings.Contains(req.Email, "@") {
		return errors.New("Wrong email format")
	}
	if len(req.Username) < 6 {
		return errors.New("Username should have at least 6 characters")
	}
	if len(req.Password) < 8 {
		return errors.New("Password should have at least 8 characters")
	}
	return nil
}

func ValidateUpdateUserProfileRequest(req *entity.UpdateUserRequest) error {
	if len(req.FirstName) < 1 {
		return errors.New("First name can't be empty.")
	}
	if len(req.LastName) < 1 {
		return errors.New("Last name can't be empty.")
	}
	if len(req.PhoneNumber) < 11 {
		return errors.New("Phone number should have at least 11 number")
	}
	if len(req.Location) < 2 {
		return errors.New("Location should have at least 2 characters")
	}
	if !ValidatePhoneNumber(req.PhoneNumber) {
		return errors.New("Wrong phone number format.")

	}

	return nil
}

func ValidatePhoneNumber(phoneNumber string) bool {
	pattern := `^\+[1-9]\d{1,14}$`

	regex := regexp.MustCompile(pattern)

	return regex.MatchString(phoneNumber)
}

func ValidateUpdatePasswordRequest(req *entity.UpdatePasswordRequest) error {
	if req.Password != req.ConfirmPassword {
		return errors.New("Passwrod didn't match.")
	}
	if len(req.Password) < 8 || len(req.ConfirmPassword) < 8 {
		return errors.New("Password should have at least 8 characters")
	}

	return nil
}