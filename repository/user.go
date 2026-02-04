package repository

import (
	"database/sql"
	"demo02/global"
	"demo02/model"
	"fmt"
	"log"
)

type UserDAO struct {
	db *sql.DB
}

func NewUserDAO() *UserDAO {
	return &UserDAO{
		db: global.GetDb(),
	}
}

func (u *UserDAO) CheckUserExisit(user *model.User) (bool, error) {
	var exists bool
	err := global.Db.QueryRow("select exists(select 1 from 用户 where 用户名=?)", user.Username).Scan(&exists)
	if err != nil {
		log.Fatal("error:%v", err)
		return false, err
	}
	if exists == true {
		return true, nil
	}
	err = global.Db.QueryRow("select exists(select 1 from 用户 where 手机号=?)", user.Phonenumber).Scan(&exists)
	if err != nil {
		log.Fatal("error:%v", err)
		return false, err
	}
	if exists == true {
		return true, nil
	}
	err = global.Db.QueryRow("select exists(select 1 from 用户 where 邮箱=?)", user.Email).Scan(&exists)
	if err != nil {
		log.Fatal("error:%v", err)
		return false, err
	}
	if exists == true {
		return true, nil
	}
	return exists, nil

}

func (u *UserDAO) InsertUser(user *model.User) error {
	_, err := global.Db.Exec("insert into 用户(用户名,密码,手机号,邮箱)value(?,?,?,?)", user.Username, user.Password, user.Phonenumber, user.Email)
	if err != nil {
		log.Fatal("注册失败%v", err)
		return err
	}
	return nil
}

func (u *UserDAO) GetUserByName(userName string) (*model.User, error) {
	use := new(model.User)
	user, err := global.Db.Query(`select 用户名,密码,手机号,邮箱,用户id from 用户 where 用户名 = ?`, userName)
	if err != nil {
		log.Fatal("error:%v", err)
		return nil, err
	}
	if user == nil {
		log.Fatal("error:%v", err)
		return nil, err
	}
	defer user.Close()
	for user.Next() {
		if err = user.Scan(&use.Username, &use.Password, &use.Phonenumber, &use.Email, &use.UserID); err != nil {
			log.Fatal("Scan失败", err)
			return nil, err
		}
	}
	return use, nil
}

// 根据用户ID查找该用户的信息
func (u *UserDAO) GetUserByID(userID int) (*model.User, error) {
	use := new(model.User)
	user, err := global.Db.Query(`select 用户名,密码,手机号,邮箱,用户id from 用户 where 用户id = ?`, userID)
	if err != nil {
		log.Fatal("error:%v", err)
		return nil, err
	}
	defer user.Close()
	for user.Next() {

		if err = user.Scan(&use.Username, &use.Password, &use.Phonenumber, &use.Email, &use.UserID); err != nil {
			log.Fatal("Scan失败", err)
			return nil, err
		}
	}
	return use, nil
}

// 更改用户数据
func (u *UserDAO) UpdateUserInfo(username string, email string, phonenumber string, userID int) error {
	res, err := global.Db.Exec("UPDATE user SET 用户名=?,邮箱=?,手机号=?, WHERE 用户id=?", username, email, phonenumber, userID)
	if err != nil {
		log.Fatal("error:%v", err)
		return err
	}
	// 3. 获取受影响行数（可选，验证是否更新成功）
	rowsAffected, _ := res.RowsAffected()
	fmt.Printf("成功更新 %d 条数据\n", rowsAffected)
	return nil
}

func (u *UserDAO) ChangePassword(newpassword string, userID int) error {
	_, err := global.Db.Exec("UPDATE user SET 密码=? WHERE 用户id=?", newpassword, userID)
	if err != nil {
		log.Fatal("error:%v", err)
		return err
	}
	return nil
}
