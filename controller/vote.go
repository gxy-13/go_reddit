package controller

import (
	"fmt"
	"go_reddit/dao/redis"
	"go_reddit/models"

	"github.com/gin-gonic/gin"
)

func VoteController(c *gin.Context) {
	// 给哪个文章点赞还是踩
	var vote models.VoteData
	if err := c.ShouldBindJSON(&vote); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		return
	}
	userID, err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	if err := redis.PostVote(vote.PostID, fmt.Sprint(userID), vote.Direction); err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
