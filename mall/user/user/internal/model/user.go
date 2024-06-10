package model

type User struct {
	Name   string
	Gender string
	Id     int64
}

func (u *User) TableName() string {
	return "user"
}
