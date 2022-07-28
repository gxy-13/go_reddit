package controller

import (
	"fmt"
	"go_reddit/dao/mysql"
	"go_reddit/models"
	"go_reddit/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// PostController 发布controller
func PostController(c *gin.Context) {
	fmt.Println("post controller .....")
	//获取输入的内容
	var p models.Post
	if err := c.ShouldBindJSON(&p); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		return
	}
	fmt.Println(p)
	// 获取当前用户id
	userID, err := GetCurrentUser(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID() failed", zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID

	err = service.CreatePost(&p)
	if err != nil {
		zap.L().Error("service.Post() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	fmt.Printf("controller ====:%v\n", p)
	ResponseSuccess(c, nil)
}

// PostListController 获取所有post
func PostListController(c *gin.Context) {
	posts, err := mysql.GetPostList()
	if err != nil {
		zap.L().Error("mysql.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	for _, v := range posts {
		fmt.Println(v.PostID)
	}
	ResponseSuccess(c, posts)
}

// PostDetailController 获取Post详细信息
func PostDetailController(c *gin.Context) {
	// 获取post id
	id := c.Param("id")
	postID, err := strconv.ParseUint(id, 10, 64)
	fmt.Printf("controller ------ id:%d\n", postID)
	post, err := service.GetPost(postID)
	if err != nil {
		zap.L().Error("service.GetPost(postID) failed", zap.Error(err))
	}
	ResponseSuccess(c, post)
}
