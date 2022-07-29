package service

import (
	"go_reddit/dao/mysql"
	"go_reddit/models"
)

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
