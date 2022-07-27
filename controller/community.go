package controller

import (
	"go_reddit/dao/mysql"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// CommunityController 社区Controller
func CommunityController(c *gin.Context) {
	// 调用service 查询所有社区（id， name）
	communityList, err := mysql.GetCommunityList()
	if err != nil {
		zap.L().Error("mysql.GetCommunityList failed.", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}
	ResponseSuccess(c, communityList)
}
