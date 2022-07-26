package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"go_reddit/models"
)

const secret = "gxy"

var (
	ErrorUserExist     = errors.New("用户已存在")
	ErrorUserNotExist  = errors.New("用户不存在")
	ErrorWrongPassword = errors.New("密码错误")
)

// CheckUserExist 检查用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser 向数据库中插入新的用户
func InsertUser(user *models.User) (err error) {
	fmt.Println("dao -----")
	fmt.Println(user)
	// 对密码进行加密
	user.Password = encryptPassword(user.Password)
	fmt.Println(user.Password)
	// 执行sql语句
	sqlStr := `insert into user(user_id,username,password) values(?,?,?)`
	db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}

func CheckLogin(u *models.User) (err error) {
	// 保存一下用户的旧密码
	oPassword := u.Password
	fmt.Println(u.Username)
	sqlStr := `select user_id, username, password from user where username = ?`
	err = db.Get(u, sqlStr, u.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	// 判断密码是否正确
	password := encryptPassword(oPassword)
	fmt.Println(password)
	fmt.Println(u.Password)
	if password != u.Password {
		return ErrorWrongPassword
	}
	return
}
