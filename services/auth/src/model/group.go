package model

type Group struct {
	BaseModel `gorm:"embedded"`
	Name      string `gorm:"not null"`
}
