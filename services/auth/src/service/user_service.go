package service

import (
	"github.com/A-Siam/bracker/auth/src/dto"
	"github.com/A-Siam/bracker/auth/src/system_error"
)

func CreateUser(createUserDto dto.CreateUserDto) (dto.UserDto, system_error.LogicalError) {
	return dto.UserDto{}, nil
}
