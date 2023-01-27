package model

type User struct {
	ID       string  `gorm:"primaryKey"`
	Name     string  `gorm:"not null"`
	Username string  `gorm:"not null;unique"`
	Password string  `gorm:"not null"`
	Groups   []Group `gorm:"many2many:user_groups;"`
}
