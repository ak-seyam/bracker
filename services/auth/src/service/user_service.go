package service

import (
	"net/http"

	"github.com/A-Siam/bracker/auth/src/common/loggers"
	"github.com/A-Siam/bracker/auth/src/common/system_errors"
	"github.com/A-Siam/bracker/auth/src/dto"
	"github.com/A-Siam/bracker/auth/src/message"
	"github.com/A-Siam/bracker/auth/src/model"
	"github.com/A-Siam/bracker/auth/src/model/dal"
	"github.com/A-Siam/bracker/auth/src/service/password_services"
	"github.com/A-Siam/bracker/auth/src/service/tokens_services"
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
	producerErr := message.ProduceNewUserMessage(userDto)
	if producerErr != nil {
		loggers.WarningLogger.Println("Failed to send the message [reason:", producerErr.Error(), "]. Please consider retrying!")
	}
	return userDto, nil
}

func Login(userLoginDto dto.UserLoginDto) (dto.UserLoginResponseDto, error) {
	user, err := dal.FindUserByUsername(userLoginDto.Username)
	if err != nil {
		return dto.UserLoginResponseDto{}, err
	}
	isPasswordCorrect, err := password_services.NewArgon2Identifier().Verify(userLoginDto.Password, user.Password)
	if err != nil {
		return dto.UserLoginResponseDto{}, err
	}
	if !isPasswordCorrect {
		return dto.UserLoginResponseDto{}, system_errors.NewLogicalError("Invalid password", http.StatusBadRequest)
	}
	return generateTokens(user)
}

func generateTokens(user model.User) (dto.UserLoginResponseDto, error) {
	accessToken, err := tokens_services.GetAccessToken(tokens_services.ConvertUserToClaims(user, 5))
	if err != nil {
		return dto.UserLoginResponseDto{}, err
	}
	refreshToken, err := tokens_services.GetRefreshToken(tokens_services.ConvertUserToClaims(user, 120))
	if err != nil {
		return dto.UserLoginResponseDto{}, err
	}
	return dto.UserLoginResponseDto{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
