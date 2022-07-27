package controller

import (
	"errors"
	"fmt"
	"go_reddit/dao/mysql"
	"go_reddit/models"
	"go_reddit/pkg/jwt"
	"go_reddit/service"
	"net/http"
	"strings"

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

	ResponseSuccess(c, CodeSuccess)
}

// SignInController 用户登录controller
func SignInController(c *gin.Context) {
	// 从请求中获取参数
	// 创建用户实例
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		zap.L().Error("Sign in with invalid param", zap.Error(err))
		//c.JSON(http.StatusOK, gin.H{
		//	"message": err.Error(),
		//})
		ResponseError(c, CodeInvalidParam)
		return
	}
	if err := mysql.CheckLogin(&user); err != nil {
		zap.L().Error("mysql.Login(&u) failed", zap.Error(err))
		ResponseError(c, CodeWrongPassword)
	}
	// 生成Token
	aToken, rToken, _ := jwt.GenToken(user.UserID)
	ResponseSuccess(c, gin.H{
		"accessToken":  aToken,
		"refreshToken": rToken,
		"userID":       user.UserID,
		"username":     user.Username,
	})
}
func RefreshTokenController(c *gin.Context) {
	rt := c.Query("refresh_token")
	// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
	// 这里假设Token放在Header的Authorization中，并使用Bearer开头
	// 这里的具体实现方式要依据你的实际业务情况决定
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		ResponseErrorWithMsg(c, CodeInvalidToken, "请求头缺少Auth Token")
		c.Abort()
		return
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		ResponseErrorWithMsg(c, CodeInvalidToken, "Token格式不对")
		c.Abort()
		return
	}
	aToken, rToken, err := jwt.RefreshToken(parts[1], rt)
	fmt.Println(err)
	c.JSON(http.StatusOK, gin.H{
		"access_token":  aToken,
		"refresh_token": rToken,
	})
}
