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
	page, size := getPageInfo(c)
	posts, err := mysql.GetPostList(page, size)
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

// GetPostListController2升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口(api分组展示使用的)
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts2 [get]
func GetPostListController2(c *gin.Context) {
	// GET请求参数(query string)：/api/v1/posts2?page=1&size=10&order=time
	// 初始化结构体时指定初始参数
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListController2 with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 合二为一
	data, err := service.GetPostListNew(p)
	// 获取数据
	if err != nil {
		zap.L().Error("service.GetPostListNew() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
