package controller

import (
	"errors"
	"fmt"
	"go_reddit/dao/mysql"
	"go_reddit/models"
	"go_reddit/service"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// SignUpController 用户注册controller
func SignUpController(c *gin.Context) {
	// 1. 参数校验
	var p models.ParamSignUp
	// 校验数据格式
	if err := c.ShouldBindJSON(&p); err != nil {
		// 请求参数出错
		zap.L().Error("Sign Up with invalid param", zap.Error(err))
		// 判断err是不是validator.validationErrors 类型
		//errs, ok := err.(validator.ValidationErrors)
		//if !ok {
		//	c.JSON(http.StatusOK, gin.H{
		//		"message": err.Error(),
		//	})
		//	return
		//}

		//c.JSON(http.StatusOK, gin.H{
		//	"message": err.Error(), // 翻译错误
		//})
		ResponseError(c, CodeInvalidParam)

		return
	}
	// 手动进行参数校验
	//if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.RePassword != p.Password {
	//	zap.L().Error("Sign Up with invalid param")
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "wrong param",
	//	})
	//	return
	//}

	fmt.Println(p)
	// 2. 业务处理
	if err := service.SignUp(&p); err != nil {
		zap.L().Error("signup failed", zap.Error(err))
		//c.JSON(http.StatusOK, gin.H{
		//	"message": "注册失败",
		//})
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "success",
	//})
	ResponseSuccess(c, CodeSuccess)
}

// SignInController 用户登录controller
func SignInController(c *gin.Context) {
	// 从请求中获取参数
	// 创建用户实例
	var user models.Login
	if err := c.ShouldBindJSON(&user); err != nil {
		zap.L().Error("Sign in with invalid param", zap.Error(err))
		//c.JSON(http.StatusOK, gin.H{
		//	"message": err.Error(),
		//})
		ResponseError(c, CodeInvalidParam)
		return
	}
	fmt.Println(user)
	// 交给service层
	token, err := service.SignIn(&user)
	if err != nil {
		zap.L().Error("service.Login failed", zap.String("username", user.Username))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeWrongPassword)
		return
	}
	ResponseSuccess(c, token)
}
