package mysql

import (
	"database/sql"
	"fmt"
	"go_reddit/models"
	"strings"

	"github.com/jmoiron/sqlx"

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

func GetPostList(page, size int64) (posts []*models.ApiPostDetail, err error) {
	sqlStr := "select post_id, title, content, author_id, community_id, create_time from post order by  create_time Desc limit ?, ?"
	posts = make([]*models.ApiPostDetail, 0, 2)
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}

func GetPostByID(postID uint64) (post *models.Post, err error) {
	fmt.Printf("get post by id: %d\n", postID)
	post = new(models.Post)
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

// GetPostListByIDs 根据给定id列表查询帖子数据
func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time from post where post_id in (?) order by FIND_IN_SET(post_id,?)`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	return
}
