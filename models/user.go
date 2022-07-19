package models

type User struct {
	Id         string `json:"id" gorm:"PRIMARY_KEY;NOT_NULL"`
	Pw         string `json:"password"`
	Name       string `json:"name"`
	Isstudent  string `json:"isStudent"`
	Department string `json:"department"`
	Age        int    `json:"age"`
	Cardinal   string `json:"cardinal"`

	UserBoard []Board `gorm:"foreignKey:UserId;association_foreignkey:Id"`
}

type Board struct {
	Num     int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT_NULL"`
	UserId  string `json:"userid" gorm:"NOT_NULL"`
	Content string `json:"contents" gorm:"NOT_NULL"`
}