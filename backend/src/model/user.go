package model

type User struct {
	Id       int    `gorm:"primaryKey;not null;autoIncrement" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}

func (User) TableName() string {
	return "user"
}
