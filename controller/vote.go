package controller

import (
	"fmt"
	"go_reddit/models"
	"go_reddit/service"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func VoteController(c *gin.Context) {
	// 给哪个文章点赞还是踩
	vote := new(models.VoteData)
	if err := c.ShouldBindJSON(vote); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		return
	}
	fmt.Println(vote)
	userID, err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	// 具体投票的业务逻辑
	if err := service.VoteForPost(userID, vote); err != nil {
		zap.L().Error("service.VoteForPost() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//if err := redis.PostVote(vote.PostID, fmt.Sprint(userID), vote.Direction); err != nil {
	//	ResponseError(c, CodeServerBusy)
	//	return
	//}
	ResponseSuccess(c, nil)
}
