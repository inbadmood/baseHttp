package _struct

type User struct {
	ID     int    `gorm:"column:id"`
	Name   string `gorm:"column:name"`
	Serial string `gorm:"column:serial"`
}
