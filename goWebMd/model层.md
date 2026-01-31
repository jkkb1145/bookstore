# model层基本功能

定义结构体,存放网站用户,书籍等数据的信息

一个结构体对应数据库中的一张表,表中的列需要包括结构体中的全部元素

###1.User

```go
type User struct {
	UserID      int    `json:"id"`
	Username    string `form:"username" json:"username"`
	Password    string `form:"password" json:"password"`
	Phonenumber string `form:"phonenum" json:"phonenum"`
	Email       string `form:"email" json:"email"`
	IsAdmin     int    `gorm:"default:false" json:"is_admin"` // 是否为管理员
}
```

User是一个存放用户数据的结构体

相应的，数据库中需要建立一张用户表，表中包括用户id,用户名,密码,手机号,邮箱,是否为管理员六列