package controller

import (
	"demo02/initjwt"
	"demo02/model"
	"demo02/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// controler层调用service层的方法，所以controller层里要有service的成员
type UserControl struct {
	//UserService 就是一个service层的结构体变量
	UserService *service.UserService
}

// 实例化UserControl
func NewUserCTL() *UserControl {
	return &UserControl{
		UserService: service.NewUserSVC(),
	}
}

//以下为供路由层调用的方法

// RegisterRequest 注册请求结构体
// 访问注册用户的路由时，附带的JSON数据要包括该结构体中所有元素
type RegisterRequest struct {
	Username     string `json:"username" binding:"required"`
	Password     string `json:"password" binding:"required"`
	ConfirmPaswd string `json:"confirmpaswd" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	Phonenumber  string `json:"phonenumber" binding:"required"`
}

// 注册一个用户
func (u *UserControl) UserRegistIn(ctx *gin.Context) {
	var user RegisterRequest
	//shouldbind将前端传来的JSON表单数据绑定给user
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(500, gin.H{
			"code":  -1,
			"msg":   "参数绑定失败",
			"error": err,
		})
		return
	}

	//在前端显示user绑定的数据，检验是否一致
	ctx.JSON(http.StatusOK, gin.H{
		"userName":    user.Username,
		"passWord":    user.Password,
		"phoneNum":    user.Phonenumber,
		"email":       user.Email,
		"cfmPassword": user.ConfirmPaswd,
	})

	//检验两次密码是否一致
	if user.Password != user.ConfirmPaswd {
		ctx.JSON(400, gin.H{
			"code": -1,
			"msg":  "两次密码不一致",
		})
		return
	}

	//调用service层方法
	//UserCTL是一个存放了UserService的结构体，可以通过调用UserCTL来调用service层的方法
	reu := model.User{
		Username:    user.Username,
		Password:    user.Password,
		Phonenumber: user.Phonenumber,
		Email:       user.Email,
	}
	if err := u.UserService.UserRegist(&reu); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  "用户注册失败",
		})
		return
	}

}

// 已注册的用户登录
func (u *UserControl) UserLogIn(ctx *gin.Context) {
	var loginuser model.LogInU

	if err := ctx.ShouldBindBodyWithJSON(&loginuser); err != nil {
		ctx.JSON(500, gin.H{
			"code":  -1,
			"msg":   "登录参数绑定失败",
			"error": err.Error(),
		})
		return
	}

	response, err := u.UserService.UserLogIn(loginuser.Username, loginuser.Password)
	if err != nil {
		ctx.JSON(500, gin.H{
			"code":  -1,
			"msg":   "登陆失败",
			"error": err.Error(),
		})
		return
	} else {
		ctx.JSON(200, gin.H{
			"code": 0,
			"data": response,
			"msg":  "登陆成功",
		})
	}
}

//TODO: 1.为用户的密码加密 2.使用JWT用于验证唯一用户，连接redis将JWT用于验证用户的token存放在内存中 3.用户登出，删除redis中相关token

func (u *UserControl) GetUserProfile(ctx *gin.Context) {
	userID, exists := ctx.Get("UserID")
	if !exists {
		//没有找到响应ID的用户，返回一个用户不存在的信息
		ctx.JSON(401, gin.H{
			"coded": -1,
			"msg":   "未登录",
		})
		return
	}
	user, err := u.UserService.GetUserByID(userID.(int))
	if err != nil {
		ctx.JSON(500, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	//存储用户的相关信息
	response := gin.H{
		"id":       user.UserID,
		"username": user.Username,
		"email":    user.Email,
		"phone":    user.Phonenumber,
		"is_admin": user.IsAdmin,
	}
	//返回获取成功的信息
	ctx.JSON(200, gin.H{
		"code": 0,
		"data": response,
		"msg":  "获取用户信息成功",
	})
}

// 更改部分用户信息，无更改密码
func (u *UserControl) UpdateUserProfile(ctx *gin.Context) {
	userID, exists := ctx.Get("UserID")
	if !exists {
		ctx.JSON(401, gin.H{
			"coded": -1,
			"msg":   "未登录",
		})
	}
	//更新后的用户信息
	var updateDate struct {
		Username    string `json:"username"`
		Email       string `json:"email"`
		Phonenumber string `json:"phonenumber"`
	}
	if err := ctx.ShouldBindJSON(&updateDate); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":  -1,
			"msg":   "绑定失败",
			"error": err.Error(),
		})
		return
	}
	//转到后端层
	err := u.UserService.UpdateUserInfo(updateDate.Username, updateDate.Email, updateDate.Phonenumber, userID.(int))
	if err != nil {
		ctx.JSON(500, gin.H{
			"code":  -1,
			"msg":   "用户信息更新失败",
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg":  "用户信息更新成功",
	})
}
func (u *UserControl) ChangePassword(c *gin.Context) {
	userID, exists := c.Get("UserID")
	if !exists {
		c.JSON(401, gin.H{
			"code": -1,
			"msg":  "未登录",
		})
	}
	//修改密码所需的信息：旧密码与新密码
	var passwordData struct {
		OldpPassword string `json:"old_password"`
		NewPassword  string `json:"new_password"`
	}
	if err := c.ShouldBindJSON(&passwordData); err != nil {
		c.JSON(500, gin.H{
			"code": -1,
			"msg":  "绑定失败",
		})
	}
	err := u.UserService.ChangePassword(passwordData.OldpPassword, passwordData.NewPassword, userID.(int))
	if err != nil {
		c.JSON(500, gin.H{
			"code":  -1,
			"msg":   "修改密码失败",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "修改密码成功",
	})
}

func (u *UserControl) Logout(c *gin.Context) {
	userID, exists := c.Get("UserID")
	if !exists {
		c.JSON(401, gin.H{
			"code": -1,
			"msg":  "未登录",
		})
		return
	}
	//退出登录后，需删除该用户的JWTtoken，调用RevikeToken实现
	err := initjwt.RevokeToken(uint(userID.(int)))
	if err != nil {
		c.JSON(500, gin.H{
			"code": -1,
			"msg":  "登出失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "用户已登出",
	})
}
