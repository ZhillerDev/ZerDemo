package model

type User struct {
	Id       int `gorm:"AUTO_INCREMENT;not null;primary_key"`
	Username string
	Age      int
	Email    string
	AddTime  int
}

func (User) TableName() string {
	return "user"
}
