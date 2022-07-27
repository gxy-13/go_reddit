package service

import (
	"fmt"
	"go_reddit/dao/mysql"
	"go_reddit/dao/redis"
	"go_reddit/models"
	"go_reddit/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	// 生成帖子id
	snowflake.Init(1)
	postID, err := snowflake.GetID()
	if err != nil {
		zap.L().Error("snowflake.GetID() failed", zap.Error(err))
		return
	}
	p.PostID = postID
	// 创建帖子
	if err := mysql.CreatePost(p); err != nil {
		zap.L().Error("mysql.CreatePost() failed", zap.Error(err))
		return err
	}
	community, err := mysql.GetCommunityNameByID(fmt.Sprint(p.CommunityID))
	if err != nil {
		zap.L().Error("mysql.GetCommunityNameByID failed", zap.Error(err))
		return err
	}
	if err := redis.CreatePost(
		fmt.Sprint(p.PostID),
		fmt.Sprint(p.AuthorID),
		p.Title,
		TruncateByWords(p.Content, 120),
		community.CommunityName); err != nil {
		zap.L().Error("redis.CreatePost failed", zap.Error(err))
		return err
	}
	return
}

func GetPost(postID string) (post *models.ApiPostDetail, err error) {
	post, err = mysql.GetPostByID(postID)
	if err != nil {
		zap.L().Error("mysql.GetPostByID() failed", zap.Error(err))
		return nil, err
	}
	user, err := mysql.GetUserByID(fmt.Sprint(post.AuthorID))
	if err != nil {
		zap.L().Error("mysql.GetUserByID() failed", zap.Error(err))
		return
	}
	post.AuthorName = user.Username
	community, err := mysql.GetCommunityByID(fmt.Sprint(post.CommunityID))
	if err != nil {
		zap.L().Error("mysql.GetCommunityByID() failed", zap.Error(err))
		return
	}
	post.CommunityName = community.CommunityName
	return post, nil
}
