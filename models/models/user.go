package models

type User struct {
	UserId   int    `orm:"column(user_id);pk"`
	Password string `orm:"column(password);size(128)"`
	UserName string `orm:"column(user_name);size(128)"`
	Phone    string `orm:"column(phone);size(11)"`
	Image    string `orm:"column(image);size(128)"`
}


func (t *User) TableName() string {
	return "article"
}