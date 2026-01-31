package model

type User struct {
	UserID      int    `json:"id"`
	Username    string `form:"username" json:"username"`
	Password    string `form:"password" json:"password"`
	Phonenumber string `form:"phonenum" json:"phonenum"`
	Email       string `form:"email" json:"email"`
	IsAdmin     int    `gorm:"default:false" json:"is_admin"` // 是否为管理员
}

type LogInU struct {
	Username string `form:"username" ,json:"username"`
	Password string `form:"password" ,json:"password"`
}
