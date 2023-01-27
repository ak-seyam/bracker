package validator

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/A-Siam/bracker/auth/src/common/system_errors"
	"github.com/A-Siam/bracker/auth/src/dto"
)

func ValidateUser(user dto.CreateUserDto) *system_errors.LogicalError {
	if len(user.Name) < 3 {
		return system_errors.NewLogicalError("username should be greater than 3", http.StatusBadRequest)
	}
	if err := passwordValidation(user.Password); err != nil {
		return system_errors.NewLogicalError(err.Error(), http.StatusBadRequest)
	}
	return nil
}

func passwordValidation(password string) error {
	if len(password) < 8 {
		return errors.New("you password must be 8 characters or more")
	}
	if hasLowercaseCharacter, err := regexp.MatchString("[a-z]+", password); err != nil {
		return err
	} else if !hasLowercaseCharacter {
		return errors.New("your password must have at least a lowercase character")
	}
	if hasUppercaseCharacter, err := regexp.MatchString("[A-Z]+", password); err != nil {
		return err
	} else if !hasUppercaseCharacter {
		return errors.New("you password must have at least one uppercase character")
	}
	if hasNumber, err := regexp.MatchString("[0-9]+", password); err != nil {
		return err
	} else if !hasNumber {
		return errors.New("your password must have at least one number")
	}
	if hasSpecialCharacter, err := regexp.MatchString("^\\w+", password); err != nil {
		return err
	} else if !hasSpecialCharacter {
		return errors.New("your password must have at least one special character")
	}
	return nil
}
