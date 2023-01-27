package model

type Group struct {
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}
