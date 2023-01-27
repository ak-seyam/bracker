package model

type User struct {
	BaseModel `gorm:"embedded"`
	Name      string  `gorm:"not null"`
	Username  string  `gorm:"not null;unique"`
	Password  string  `gorm:"not null"`
	Groups    []Group `gorm:"many2many:user_groups;"`
}

type UserGroup struct {
	UserId  string `gorm:"primaryKey"`
	GroupId string `gorm:"primaryKey"`
}
