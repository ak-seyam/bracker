package dal_mappers

import (
	"github.com/A-Siam/bracker/auth/src/dto"
	"github.com/A-Siam/bracker/auth/src/model"
	"github.com/A-Siam/bracker/auth/src/service/password_services"
)

func MapCreateUserDtoToModel(id string, createUserDto dto.CreateUserDto) (model.User, error) {
	aid := password_services.NewArgon2Identifier()
	password, err := aid.HashPassword(createUserDto.Password)
	if err != nil {
		return model.User{}, err
	}
	return model.User{
		Name:     createUserDto.Name,
		Password: password,
		BaseModel: model.BaseModel{
			ID: id,
		},
		Username: createUserDto.Username,
	}, nil
}

func MapModelToUserDto(model model.User) dto.UserDto {
	return dto.UserDto{
		ID:       model.ID,
		Username: model.Username,
		Name:     model.Name,
		Groups:   getGroups(model.Groups),
	}
}

func getGroups(groups []model.Group) []string {
	res := []string{}
	for _, group := range groups {
		res = append(res, group.Name)
	}
	return res
}
