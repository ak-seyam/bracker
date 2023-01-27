package dal

import (
	"github.com/A-Siam/bracker/auth/src/constants"
	"github.com/A-Siam/bracker/auth/src/database"
	"github.com/A-Siam/bracker/auth/src/model"
)

func GetDefaultGroup() (model.Group, error) {
	db, err := database.GetDB()
	if err != nil {
		return model.Group{}, err
	}
	var group = &model.Group{}
	tx := db.Where("name = ?", constants.DEFAULT_GROUP_NAME).Find(group)
	if tx.Error != nil {
		return model.Group{}, tx.Error
	}
	return *group, nil
}
