package mysql

import (
	"database/sql"
	"go_reddit/models"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	sqlStr := "insert into post(post_id, title, content, author_id, community_id) values(?,?,?,?,?)"
	_, err = db.Exec(sqlStr, p.PostID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	if err != nil {
		zap.L().Error("insert into post failed", zap.Error(err))
		return
	}
	return
}

func GetPostList() (posts []*models.ApiPostDetail, err error) {
	sqlStr := "select post_id, title, content, author_id, community_id, create_time from post limit 2"
	posts = make([]*models.ApiPostDetail, 0, 2)
	err = db.Select(&posts, sqlStr)
	return
}

func GetPostByID(postID string) (post *models.ApiPostDetail, err error) {
	sqlStr := "select post_id, title, content, author_id, community_id, create_time from post where post_id = ?"
	err = db.Get(post, sqlStr, postID)
	if err == sql.ErrNoRows {
		zap.L().Error("postId error", zap.Error(err))
		return
	}
	if err != nil {
		zap.L().Error("query post failed", zap.Error(err))
		return
	}
	return
}
