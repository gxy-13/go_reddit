package controller

import (
	"go_reddit/dao/mysql"
	"go_reddit/service"
	"strconv"

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

// CommunityDetailController
func CommunityDetailController(c *gin.Context) {
	// 获取社区id
	idStr := c.Param("id") // 获取url参数
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 根据id获取社区详情
	data, err := service.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("service.GetCommunityDetail() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
