package redis

/*
	Redis Key
*/

const (
	Prefix             = "bluebell:"  //项目key前缀
	KeyPostTimeZSet    = "post:time"  // ZSet 帖子及发帖时间
	KeyPostScoreZSet   = "post:score" // ZSet 帖子及投票的分数
	KeyPostVotedZSetPF = "post:voted" // ZSet 记录用户及投票类型，参数是post id
	KeyCommunitySetPF  = "community:" // set; 保存每个分区下帖子的id
)

// 给redis key 加上前缀
func getRedisKey(key string) string {
	return Prefix + key
}
