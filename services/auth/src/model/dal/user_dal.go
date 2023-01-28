package dal

import (
	"net/http"

	"github.com/A-Siam/bracker/auth/src/common/system_errors"
	"github.com/A-Siam/bracker/auth/src/database"
	"github.com/A-Siam/bracker/auth/src/dto"
	"github.com/A-Siam/bracker/auth/src/model"
	"github.com/A-Siam/bracker/auth/src/model/dal/dal_mappers"
)

func SaveUser(id string, createUserDto dto.CreateUserDto) (dto.UserDto, error) {
	user, err := dal_mappers.MapCreateUserDtoToModel(id, createUserDto)
	if err != nil {
		return dto.UserDto{}, err
	}
	defaultGroup, err := GetDefaultGroup()
	if err != nil {
		return dto.UserDto{}, err
	}
	user.Groups = []model.Group{defaultGroup}
	db, err := database.GetDB()
	if err != nil {
		return dto.UserDto{}, err
	}
	db.Create(&user)
	tx := db.Save(&user)
	if tx.Error != nil {
		return dto.UserDto{}, tx.Error
	}
	return dal_mappers.MapModelToUserDto(user), nil
}

func FindUserByUsername(username string) (model.User, error) {
	db, err := database.GetDB()
	if err != nil {
		return model.User{}, err
	}
	user := model.User{}
	tx := db.Where("username = ?", username).Preload("Groups").Find(&user)
	if tx.Error != nil {
		return model.User{}, system_errors.NewLogicalError(tx.Error.Error(), http.StatusBadRequest)
	}
	return user, nil
}
