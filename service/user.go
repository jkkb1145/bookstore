package service

import (
	"demo02/initjwt"
	"demo02/model"
	"demo02/repository"
	"errors"
	"fmt"
)

type UserService struct {
	UserDB *repository.UserDAO
}

func NewUserSVC() *UserService {
	return &UserService{
		UserDB: repository.NewUserDAO(),
	}
}

type LogInReponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpireIn     int64     `json:"expire_in"`
	UserInfo     *UserInfo `json:"user_info"`
}

type UserInfo struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func (u *UserService) UserRegist(user *model.User) error {
	//1,防止某些数据重复
	exists, err := u.UserDB.CheckUserExisit(user)
	if err != nil {
		return err
	}
	if exists == true {
		return err
	}
	//
	if err := u.UserDB.InsertUser(user); err != nil {
		return err
	}
	return nil
}

func (u *UserService) UserLogIn(username, password string) (*LogInReponse, error) {
	//查询用户是否存在
	user, err := u.UserDB.GetUserByName(username)
	if err != nil {
		fmt.Println("获取用户信息失败")
		return nil, errors.New("获取用户信息失败")
	}
	//校验密码
	if !u.CheeckPassword(password, user.Password) {
		fmt.Println("密码错误")
		return nil, errors.New("密码错误")
	}
	//jwt校验用户
	token, err := initjwt.GenerateTokenPair(uint(user.UserID), user.Username)
	if err != nil {
		fmt.Println("token生成失败")
		return nil, errors.New("token生成失败")
	}
	reponse := &LogInReponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpireIn:     token.ExpiresIn,
		UserInfo: &UserInfo{
			ID:       user.UserID,
			Username: user.Username,
			Email:    user.Email,
			Phone:    user.Phonenumber,
		},
	}
	return reponse, nil
}

func (u *UserService) CheeckPassword(inputPassWord, truePassWord string) bool {
	return inputPassWord == truePassWord
}

func (u *UserService) GetUserByID(userID int) (*model.User, error) {
	user, err := u.UserDB.GetUserByID(userID)
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	return user, nil
}

func (u *UserService) UpdateUserInfo(username, email, phonenumber string, userID int) error {
	err := u.UserDB.UpdateUserInfo(username, email, phonenumber, userID)
	if err != nil {
		return errors.New("更新用户信息失败")
	}
	return nil
}
func (u *UserService) ChangePassword(oldPassword, newpassword string, userID int) error {
	user, err := u.UserDB.GetUserByID(userID)
	if err != nil {
		return errors.New("获取用户信息失败")
	}
	if user.Password != oldPassword {
		return errors.New("原密码错误")
	}
	err = u.UserDB.ChangePassword(newpassword, userID)
	if err != nil {
		return errors.New("修改密码失败")
	}
	return nil
}
