package service

import (
	"fmt"
	"go_reddit/dao/mysql"
	"go_reddit/models"
	"go_reddit/pkg/jwt"
	"go_reddit/pkg/snowflake"
)

// SignUp 用户注册业务层
func SignUp(p *models.ParamSignUp) (err error) {
	fmt.Println("service.SignUp ....")
	// 判断用户是否存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		// 数据库查询出错
		return err
	}
	// 生成uuid
	snowflake.Init(1)
	userID, _ := snowflake.GetID()
	fmt.Println("service=========")
	fmt.Println(userID)
	// 构造user实例，填充userid
	user := &models.User{
		userID,
		p.Username,
		p.Password,
	}
	// 持久化
	return mysql.InsertUser(user)
}

// SignIn 用户登录业务层
func SignIn(u *models.Login) (token string, err error) {
	user := &models.User{
		Username: u.Username,
		Password: u.Password,
	}
	fmt.Println("===========service")
	fmt.Println(user)
	// 判断账号密码是否正确
	if err = mysql.CheckLogin(user); err != nil {
		return "", err
	}
	// 登录成功设置token 因为传递的是user指针，所以获取userid
	return jwt.GenToken(user.UserID, user.Username)
}
