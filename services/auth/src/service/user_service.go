package service

import (
	"net/http"

	"github.com/A-Siam/bracker/auth/src/common/system_errors"
	"github.com/A-Siam/bracker/auth/src/dto"
	"github.com/A-Siam/bracker/auth/src/model/dal"
	"github.com/A-Siam/bracker/auth/src/service/validator"
	"github.com/google/uuid"
)

func CreateUser(createUserDto dto.CreateUserDto) (dto.UserDto, *system_errors.LogicalError) {
	id := uuid.NewString()
	if err := validator.ValidateUser(createUserDto); err != nil {
		return dto.UserDto{}, err
	}
	userDto, err := dal.SaveUser(id, createUserDto)
	if err != nil {
		return dto.UserDto{}, system_errors.NewLogicalError(err.Error(), http.StatusInternalServerError)
	}
	return userDto, nil
}
