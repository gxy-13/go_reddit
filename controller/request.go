package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var ErrorUserNotLogin = errors.New("用户未登录")

const UserIDKey = "userID"

// GetCurrentUser 获取当前登录的用户id
func GetCurrentUser(c *gin.Context) (userID uint64, err error) {
	uid, ok := c.Get(UserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(uint64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
